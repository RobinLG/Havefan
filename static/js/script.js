	// Load nem-browser library
	var nem = require("nem-sdk").default;

	// Create an NIS endpoint object
	var endpoint = nem.model.objects.create("endpoint")(nem.model.nodes.defaultTestnet, nem.model.nodes.defaultPort);

	// Get an empty un-prepared transfer transaction object
	var transferTransaction = nem.model.objects.get("transferTransaction");

	// Get an empty common object to hold pass and key
	var common = nem.model.objects.get("common");

	// Set default amount in view. It is text input so we can handle dot and comma as decimal mark easily (need cleaning but provided by the library)
	$("#amount").val("0");

	/**
	 * Function to update our fee in the view
	 */
	function updateFee() {
		// Check for amount errors
		if(undefined === $("#amount").val() || !nem.utils.helpers.isTextAmountValid($("#amount").val())) return alert('Invalid amount !');

		// Set the cleaned amount into transfer transaction object
		transferTransaction.amount = nem.utils.helpers.cleanTextAmount($("#amount").val());

		// Set the message into transfer transaction object
		transferTransaction.message = $("#message").val();

		// Prepare the updated transfer transaction object
		var transactionEntity = nem.model.transactions.prepare("transferTransaction")(common, transferTransaction, nem.model.network.data.testnet.id);

		// Format fee returned in prepared object
		var feeString = nem.utils.format.nemValue(transactionEntity.fee)[0] + "." + nem.utils.format.nemValue(transactionEntity.fee)[1];

		//Set fee in view
		$("#fee").html(feeString);
	}

	/**
	 * Build transaction from form data and send
	 */
	function send() {
		// Check form for errors
		if(!$("#privateKey").val() || !$("#recipient").val()) return alert('Missing parameter !');
		if(undefined === $("#amount").val() || !nem.utils.helpers.isTextAmountValid($("#amount").val())) return alert('Invalid amount !');
		if (!nem.model.address.isValid(nem.model.address.clean($("#recipient").val()))) return alert('Invalid recipent address !');

		// Set the private key in common object
		common.privateKey = $("#privateKey").val();

		// Check private key for errors
		if (common.privateKey.length !== 64 && common.privateKey.length !== 66) return alert('Invalid private key, length must be 64 or 66 characters !');
		if (!nem.utils.helpers.isHexadecimal(common.privateKey)) return alert('Private key must be hexadecimal only !');

		// Set the cleaned amount into transfer transaction object
		transferTransaction.amount = nem.utils.helpers.cleanTextAmount($("#amount").val());

		// Recipient address must be clean (no hypens: "-")
		transferTransaction.recipient = nem.model.address.clean($("#recipient").val());

		// Set message
		transferTransaction.message = $("#message").val();

		// Prepare the updated transfer transaction object
		var transactionEntity = nem.model.transactions.prepare("transferTransaction")(common, transferTransaction, nem.model.network.data.testnet.id);
		console.log("transactionEntity");
		console.log(transactionEntity);

		// Serialize transfer transaction and announce
		setTimeout(()=>{
			nem.model.transactions.send(common, transactionEntity, endpoint).then(function(res){
				// If code >= 2, it's an error
				if (res.code >= 2) {
					alert(res.message);
				} else {
					alert(res.message);
					console.log(res);
					window.close("../views/detail.html");
					window.open("../views/index.html");
				}
			}, function(err) {
				alert(err);
			});
		},3000)
	}

	// queryTxByAddrAndTxHash("TAV6OC3AUBFD73T7CFG2EQRE34D4EA355DTVJUDL", "722ec9844c23cbfc57d6d12b8f5a1378e679c20a52536770a5249d3724797eb0");
	function queryTxByAddrAndTxHash(Addr, TxHash) {
		nem.com.requests.account.transactions.all(endpoint, Addr).then(function(res) {
			console.log("\nAll transactions:");
			console.log(res);
			res.data.forEach(element => {
				console.log(element.meta.hash.data);
				if (element.meta.hash.data == TxHash) {
					var payload = element.transaction.message.payload
					console.log(payload);
					return payload;
					//console.log(CryptoJS.lib.WordArray.create(temp, uaLength););
				}
				
			});
			return res;
		}, function(err) {
			console.error(err);
		});
	}

	// On amount change we update fee in view
	$("#amount").on('change keyup paste', function() {
		updateFee();
	});

	// On message change we update fee in view
	$("#message").on('change keyup paste', function() {
		updateFee();
	});

	// Call send function when click on send button
	// $("#send").click(function() {
	// 	send();
	// 	console.log(send);
	// });

	// Initialization of fees in view
	updateFee();