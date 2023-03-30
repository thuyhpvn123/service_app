var output = document.getElementById("log-content");
var socket = new WebSocket("ws://127.0.0.1:2001/ws");
var socketActive = false;
console.log("Imported");
// * Websocket
// Connect to server successfully
var walletAddress = "aa39344b158f4004cac70bb4ace871a9b54baa1e";
const map = new Map()
var messageForm = {
  command:"get-wallet-pagination",
  limit:2,
  page:1,
};

socket.onopen = (msg) => {
  socketActive = true;


  //Send walletAddress to server
  let walletMessage = structuredClone(messageForm);
  walletMessage.type = "WalletMessage";
  walletMessage.message = walletAddress;
  // sendMessage(walletMessage);
};

// WS connection's closed
socket.onclose = (event) => {
  console.log("WS Connection is closed: ", event);
};

// WS connection having errors
socket.onerror = (error) => {
  console.log("Socket Error: ", error);
};
socket.onmessage = (msg) => {
  var data12 = JSON.parse(msg.data);
  output.innerHTML += "Server: " + msg.data + "\n";
  //switch case
}


var sendMessage = (msg) => {
  console.log(msg);
  socket.send(JSON.stringify(msg));

};
//check balance
var $getBalance = document.getElementById("get-balance");
$getBalance.addEventListener("submit", async(e) => {
  e.preventDefault();
var name;
var flag=1;
name = $('#token').val()

if( name ==''){
  flag=0
  $('.error_name').html("Please type dapp name")
}else{
  $('.error_name').html("")
}
if(flag==1){
  try{
    var inputMessage = structuredClone(messageForm);
    inputMessage.type = "GetBalance"; 
    inputMessage.message = `${name},${walletAddress}`, 
    sendMessage(inputMessage);
    console.log("getBalance")
  }catch(e){
    console.log(e)
  }
}

});
//send message
var $sendMessage = document.getElementById("send-message");
$sendMessage.addEventListener("submit", async(e) => {
  e.preventDefault();
var type;
var message;
var flag=1;
type = $('#type').val()
message = $('#message').val()

if( type ==''){
  flag=0
  $('.error_name').html("Please type type of message ")
}else{
  $('.error_name').html("")
}
if(flag==1){
  try{
    var inputMessage = structuredClone(messageForm);
    inputMessage.type = type; 
    inputMessage.message = message, 
    sendMessage(inputMessage);
    console.log("send message")
  }catch(e){
    console.log(e)
  }
}

});

//test functions
document.getElementById("test").addEventListener("click",test);

var test = () => {
  //wallet
  console.log("hello")
//
var messageForm17 = {
  command:"init-app",
  address:"fbed76bfaaad911c8459f87c7900d1e44b3cbc82",
  };
sendMessage(messageForm17);

//   var messageForm1 = {
//     command:"get-wallet-pagination",
//     limit:2,
//     page:1,
//   };
//   sendMessage(messageForm1);
// //
//   var messageForm2 = {
//     command:"get-all-wallet",
//   };
//   sendMessage(messageForm2);
// //
//   var getWalletAtAddressOj={
//   data:"b2660cd6151cc04b76cd9512e3bab5900f7491a2",
//   }
//   var messageForm3 = {
//     command:"getWalletAtAddress",
//     value: getWalletAtAddressOj,
//   };
//   sendMessage(messageForm3);

// //
// var messageForm4 = {
//   command:"download",
//   message:"Metanode",
// };
// sendMessage(messageForm4);
// //
var valueObj ={
  language:"english"
}
var messageForm14 = {
  command:"get-raw-seed",
  value: valueObj,
};
sendMessage(messageForm14);
//
var lastNodeConnected ={
  ip: "35.243.233.132",
  port: "3011"
}
var node ={
  ip: "35.243.233.132",
  port: "3011"
}

var wallets1 ={
  address: "5de5635ede4641b71d466d155a771d7578b086a4",
}
var wallets2 ={
  address: "e40844ec9eae618baed9c3e5c951e0e022869d97",
}
var wallets =[wallets1,wallets2]
var rawSeedRestore=["town","thought","pink", "catalog", "alcohol", "believe", "parent"," scene ","cart ","unlock ","weasel"," nothing ","lava ","same"," toy"," whip"," solve"," dignity"," usage"," shiver"," harbor"," satoshi ","tattoo ","topple"];
var valueObj1 ={
  bg:                  "green",
  color:               "yellow",
  idSymbol:            "bbb",
  pattern:             "special",
  'last-node-connected': lastNodeConnected,
  'raw-seed-restore':    rawSeedRestore,
  node:                node,
  wallets:             wallets,

}

var messageForm15 = {
  command:"create-wallet",
  value: valueObj1,
  };
sendMessage(messageForm15);

// // chi get duoc vi da duoc init-connection
// var messageForm5 = {
//   command:"get-wallet-info",
//   address:"5de5635ede4641b71d466d155a771d7578b086a4",
// };

// sendMessage(messageForm5);
//transfer money
var transfer ={
  'from-address':   "5de5635ede4641b71d466d155a771d7578b086a4",
  'to-address':    "BFC6af1caCf974d30D65e129654cC5153A56EcF6",
  amount:            "1",
  fee:             "1",
  tip:             "",
  message:    "",
  'receive-info':  "",
  'is-deploy':   false,
  'is-call':     false,
  name:            "",
  input:           "",
  image:           "",
  "abiData": "",
  "function-name":"",
  isOfflineMode:  false,
  feeType:   false,
  inputArray:"",
  maxGas:"",
  maxGasPriceGwei:"",
  maxTimeUse:"",
  relatedAddresses:[],
}

var messageForm24 = {
  command:"send-transaction",
  value: transfer,
  };
// sendMessage(messageForm24);



// var positionObj ={
//   weight:123,
//   height:456,
// }
// var dapp ={
//   "name":"game-Dragon" ,
//   "author":"thuydo"  ,
//   "hash":"hash" ,
//   "sign":"sign",
//   "version":"0.0.0.1",
//   "image":"image",
//   "pathStorage":"pathStorage",
//   "time":152355666 ,
//   "totalWallet":6468461666 ,
//   "totalTransaction":6546464321646,
//   "size":12345,
//   "bundleId":"bundleId" ,
//   "orientation":"orientation",
//   "urlWeb":"urlWeb",
//   "isLocal":46646,
//   "fullScreen":4648681,
//   "statusBar":"statusBar",
//   "groupId":545,
//   "isShowInApp":4662,
//   "page":5,
//   "position":684665,
//   "positionObj":positionObj,
//   "isInstalled":8522,
//   "abiData":"abiData",
//   "binData":"binData",
//   "status":3165,
//   "type":65,

// }
// var messageForm21 = {
//   command:"deploy-d-app",
//   value: dapp,
//   };
// sendMessage(messageForm21);
//cake
var abiDataCake = JSON.stringify(
  [
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "owner",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "spender",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "value",
          "type": "uint256"
        }
      ],
      "name": "Approval",
      "type": "event"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "spender",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "approve",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "spender",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "subtractedValue",
          "type": "uint256"
        }
      ],
      "name": "decreaseAllowance",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "delegatee",
          "type": "address"
        }
      ],
      "name": "delegate",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "delegatee",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "nonce",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "expiry",
          "type": "uint256"
        },
        {
          "internalType": "uint8",
          "name": "v",
          "type": "uint8"
        },
        {
          "internalType": "bytes32",
          "name": "r",
          "type": "bytes32"
        },
        {
          "internalType": "bytes32",
          "name": "s",
          "type": "bytes32"
        }
      ],
      "name": "delegateBySig",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "delegator",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "fromDelegate",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "toDelegate",
          "type": "address"
        }
      ],
      "name": "DelegateChanged",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "delegate",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "previousBalance",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "newBalance",
          "type": "uint256"
        }
      ],
      "name": "DelegateVotesChanged",
      "type": "event"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "spender",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "addedValue",
          "type": "uint256"
        }
      ],
      "name": "increaseAllowance",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "_amount",
          "type": "uint256"
        }
      ],
      "name": "mint",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "mint",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "previousOwner",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "newOwner",
          "type": "address"
        }
      ],
      "name": "OwnershipTransferred",
      "type": "event"
    },
    {
      "inputs": [],
      "name": "renounceOwnership",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "recipient",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "transfer",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "from",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "value",
          "type": "uint256"
        }
      ],
      "name": "Transfer",
      "type": "event"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "sender",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "recipient",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "transferFrom",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "newOwner",
          "type": "address"
        }
      ],
      "name": "transferOwnership",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "owner",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "spender",
          "type": "address"
        }
      ],
      "name": "allowance",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "account",
          "type": "address"
        }
      ],
      "name": "balanceOf",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        },
        {
          "internalType": "uint32",
          "name": "",
          "type": "uint32"
        }
      ],
      "name": "checkpoints",
      "outputs": [
        {
          "internalType": "uint32",
          "name": "fromBlock",
          "type": "uint32"
        },
        {
          "internalType": "uint256",
          "name": "votes",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "decimals",
      "outputs": [
        {
          "internalType": "uint8",
          "name": "",
          "type": "uint8"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "delegator",
          "type": "address"
        }
      ],
      "name": "delegates",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "DELEGATION_TYPEHASH",
      "outputs": [
        {
          "internalType": "bytes32",
          "name": "",
          "type": "bytes32"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "DOMAIN_TYPEHASH",
      "outputs": [
        {
          "internalType": "bytes32",
          "name": "",
          "type": "bytes32"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "account",
          "type": "address"
        }
      ],
      "name": "getCurrentVotes",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getOwner",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "account",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "blockNumber",
          "type": "uint256"
        }
      ],
      "name": "getPriorVotes",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "name",
      "outputs": [
        {
          "internalType": "string",
          "name": "",
          "type": "string"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "nonces",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "numCheckpoints",
      "outputs": [
        {
          "internalType": "uint32",
          "name": "",
          "type": "uint32"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "owner",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "symbol",
      "outputs": [
        {
          "internalType": "string",
          "name": "",
          "type": "string"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "totalSupply",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    }
  ]
);
var binDataCake = "60806040523480156200001157600080fd5b506040518060400160405280601181526020017f50616e63616b655377617020546f6b656e0000000000000000000000000000008152506040518060400160405280600581526020017f64756d6d790000000000000000000000000000000000000000000000000000008152506000620000906200018460201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3508160049080519060200190620001469291906200018c565b5080600590805190602001906200015f9291906200018c565b506012600660006101000a81548160ff021916908360ff160217905550505062000232565b600033905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001cf57805160ff191683800117855562000200565b8280016001018555821562000200579182015b82811115620001ff578251825591602001919060010190620001e2565b5b5090506200020f919062000213565b5090565b5b808211156200022e57600081600090555060010162000214565b5090565b61301080620002426000396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c8063782d6fe1116100f9578063a9059cbb11610097578063dd62ed3e11610071578063dd62ed3e1461091c578063e7a324dc14610994578063f1127ed8146109b2578063f2fde38b14610a27576101a9565b8063a9059cbb146107e7578063b4b5ea571461084b578063c3cda520146108a3576101a9565b80638da5cb5b116100d35780638da5cb5b1461068857806395d89b41146106bc578063a0712d681461073f578063a457c2d714610783576101a9565b8063782d6fe11461059a5780637ecebe00146105fc578063893d20e814610654576101a9565b806339509351116101665780635c19a95c116101405780635c19a95c146104965780636fcfff45146104da57806370a0823114610538578063715018a614610590576101a9565b8063395093511461037657806340c10f19146103da578063587cde1e14610428576101a9565b806306fdde03146101ae578063095ea7b31461023157806318160ddd1461029557806320606b70146102b357806323b872dd146102d1578063313ce56714610355575b600080fd5b6101b6610a6b565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101f65780820151818401526020810190506101db565b50505050905090810190601f1680156102235780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61027d6004803603604081101561024757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b0d565b60405180821515815260200191505060405180910390f35b61029d610b2b565b6040518082815260200191505060405180910390f35b6102bb610b35565b6040518082815260200191505060405180910390f35b61033d600480360360608110156102e757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b59565b60405180821515815260200191505060405180910390f35b61035d610c32565b604051808260ff16815260200191505060405180910390f35b6103c26004803603604081101561038c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610c49565b60405180821515815260200191505060405180910390f35b610426600480360360408110156103f057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610cfc565b005b61046a6004803603602081101561043e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e3d565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6104d8600480360360208110156104ac57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ea6565b005b61051c600480360360208110156104f057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610eb3565b604051808263ffffffff16815260200191505060405180910390f35b61057a6004803603602081101561054e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ed6565b6040518082815260200191505060405180910390f35b610598610f1f565b005b6105e6600480360360408110156105b057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506110a5565b6040518082815260200191505060405180910390f35b61063e6004803603602081101561061257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611466565b6040518082815260200191505060405180910390f35b61065c61147e565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61069061148d565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6106c46114b6565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156107045780820151818401526020810190506106e9565b50505050905090810190601f1680156107315780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61076b6004803603602081101561075557600080fd5b8101908080359060200190929190505050611558565b60405180821515815260200191505060405180910390f35b6107cf6004803603604081101561079957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061163c565b60405180821515815260200191505060405180910390f35b610833600480360360408110156107fd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611709565b60405180821515815260200191505060405180910390f35b61088d6004803603602081101561086157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611727565b6040518082815260200191505060405180910390f35b61091a600480360360c08110156108b957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803560ff16906020019092919080359060200190929190803590602001909291905050506117fd565b005b61097e6004803603604081101561093257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611b61565b6040518082815260200191505060405180910390f35b61099c611be8565b6040518082815260200191505060405180910390f35b610a04600480360360408110156109c857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803563ffffffff169060200190929190505050611c0c565b604051808363ffffffff1681526020018281526020019250505060405180910390f35b610a6960048036036020811015610a3d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611c4d565b005b606060048054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b035780601f10610ad857610100808354040283529160200191610b03565b820191906000526020600020905b815481529060010190602001808311610ae657829003601f168201915b5050505050905090565b6000610b21610b1a611e58565b8484611e60565b6001905092915050565b6000600354905090565b7f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a86681565b6000610b66848484612057565b610c2784610b72611e58565b610c2285604051806060016040528060288152602001612e8060289139600260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000610bd8611e58565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546123119092919063ffffffff16565b611e60565b600190509392505050565b6000600660009054906101000a900460ff16905090565b6000610cf2610c56611e58565b84610ced8560026000610c67611e58565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546123d190919063ffffffff16565b611e60565b6001905092915050565b610d04611e58565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610dc4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b610dce8282612459565b610e396000600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683612616565b5050565b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b610eb033826128b3565b50565b60096020528060005260406000206000915054906101000a900463ffffffff1681565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b610f27611e58565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610fe7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60004382106110ff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526027815260200180612f026027913960400191505060405180910390fd5b6000600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff16905060008163ffffffff16141561116c576000915050611460565b82600860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006001840363ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900463ffffffff1663ffffffff161161125657600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006001830363ffffffff1663ffffffff16815260200190815260200160002060010154915050611460565b82600860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008063ffffffff16815260200190815260200160002060000160009054906101000a900463ffffffff1663ffffffff1611156112d7576000915050611460565b6000806001830390505b8163ffffffff168163ffffffff1611156113fa576000600283830363ffffffff168161130957fe5b0482039050611316612dca565b600860008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008363ffffffff1663ffffffff1681526020019081526020016000206040518060400160405290816000820160009054906101000a900463ffffffff1663ffffffff1663ffffffff168152602001600182015481525050905086816000015163ffffffff1614156113d257806020015195505050505050611460565b86816000015163ffffffff1610156113ec578193506113f3565b6001820392505b50506112e1565b600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008363ffffffff1663ffffffff1681526020019081526020016000206001015493505050505b92915050565b600a6020528060005260406000206000915090505481565b600061148861148d565b905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060058054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561154e5780601f106115235761010080835404028352916020019161154e565b820191906000526020600020905b81548152906001019060200180831161153157829003601f168201915b5050505050905090565b6000611562611e58565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611622576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b61163361162d611e58565b83612459565b60019050919050565b60006116ff611649611e58565b846116fa85604051806060016040528060258152602001612f946025913960026000611673611e58565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546123119092919063ffffffff16565b611e60565b6001905092915050565b600061171d611716611e58565b8484612057565b6001905092915050565b600080600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff16905060008163ffffffff16116117915760006117f5565b600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006001830363ffffffff1663ffffffff168152602001908152602001600020600101545b915050919050565b60007f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866611828610a6b565b80519060200120611837612a24565b30604051602001808581526020018481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff16815260200194505050505060405160208183030381529060405280519060200120905060007fe48329057bfd03d55e49b547132e39cffd9c1820ad7b9d4c5307691425d15adf888888604051602001808581526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019450505050506040516020818303038152906040528051906020012090506000828260405160200180807f190100000000000000000000000000000000000000000000000000000000000081525060020183815260200182815260200192505050604051602081830303815290604052805190602001209050600060018288888860405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156119bb573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611a4d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180612ea86026913960400191505060405180910390fd5b600a60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190600101919050558914611af2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180612f4f6022913960400191505060405180910390fd5b87421115611b4b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180612e5a6026913960400191505060405180910390fd5b611b55818b6128b3565b50505050505050505050565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b7fe48329057bfd03d55e49b547132e39cffd9c1820ad7b9d4c5307691425d15adf81565b6008602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900463ffffffff16908060010154905082565b611c55611e58565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611d15576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611d9b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180612e346026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611ee6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180612e106024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611f6c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180612fb96022913960400191505060405180910390fd5b80600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156120dd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180612deb6025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415612163576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180612f716023913960400191505060405180910390fd5b6121cf81604051806060016040528060268152602001612f2960269139600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546123119092919063ffffffff16565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061226481600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546123d190919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b60008383111582906123be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015612383578082015181840152602081019050612368565b50505050905090810190601f1680156123b05780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b60008082840190508381101561244f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156124fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f42455032303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b612511816003546123d190919063ffffffff16565b60038190555061256981600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546123d190919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156126525750600081115b156128ae57600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614612782576000600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff1690506000808263ffffffff16116126f5576000612759565b600860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006001840363ffffffff1663ffffffff168152602001908152602001600020600101545b905060006127708483612a3190919063ffffffff16565b905061277e86848484612a7b565b5050505b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16146128ad576000600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff1690506000808263ffffffff1611612820576000612884565b600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006001840363ffffffff1663ffffffff168152602001908152602001600020600101545b9050600061289b84836123d190919063ffffffff16565b90506128a985848484612a7b565b5050505b5b505050565b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600061292284610ed6565b905082600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f60405160405180910390a4612a1e828483612616565b50505050565b6000804690508091505090565b6000612a7383836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250612311565b905092915050565b6000612a9f43604051806060016040528060348152602001612ece60349139612d0f565b905060008463ffffffff16118015612b3457508063ffffffff16600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006001870363ffffffff1663ffffffff16815260200190815260200160002060000160009054906101000a900463ffffffff1663ffffffff16145b15612ba55781600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006001870363ffffffff1663ffffffff16815260200190815260200160002060010181905550612cb2565b60405180604001604052808263ffffffff16815260200183815250600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008663ffffffff1663ffffffff16815260200190815260200160002060008201518160000160006101000a81548163ffffffff021916908363ffffffff1602179055506020820151816001015590505060018401600960008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff1602179055505b8473ffffffffffffffffffffffffffffffffffffffff167fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a7248484604051808381526020018281526020019250505060405180910390a25050505050565b600064010000000083108290612dc0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015612d85578082015181840152602081019050612d6a565b50505050905090810190601f168015612db25780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5082905092915050565b6040518060400160405280600063ffffffff16815260200160008152509056fe42455032303a207472616e736665722066726f6d20746865207a65726f206164647265737342455032303a20617070726f76652066726f6d20746865207a65726f20616464726573734f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737343414b453a3a64656c656761746542795369673a207369676e6174757265206578706972656442455032303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636543414b453a3a64656c656761746542795369673a20696e76616c6964207369676e617475726543414b453a3a5f7772697465436865636b706f696e743a20626c6f636b206e756d6265722065786365656473203332206269747343414b453a3a6765745072696f72566f7465733a206e6f74207965742064657465726d696e656442455032303a207472616e7366657220616d6f756e7420657863656564732062616c616e636543414b453a3a64656c656761746542795369673a20696e76616c6964206e6f6e636542455032303a207472616e7366657220746f20746865207a65726f206164647265737342455032303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726f42455032303a20617070726f766520746f20746865207a65726f2061646472657373a26469706673582212202cfd637967a463279dd86cf3b7072a6a8f6c8355135e97aeec389f4fdd5b71f364736f6c634300060c0033"
//wbnb
var abiDataWbnb = JSON.stringify([
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "src",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "guy",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "wad",
				"type": "uint256"
			}
		],
		"name": "Approval",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "dst",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "wad",
				"type": "uint256"
			}
		],
		"name": "Deposit",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "src",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "dst",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "wad",
				"type": "uint256"
			}
		],
		"name": "Transfer",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "src",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "wad",
				"type": "uint256"
			}
		],
		"name": "Withdrawal",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"name": "allowance",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "guy",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "wad",
				"type": "uint256"
			}
		],
		"name": "approve",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"name": "balanceOf",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "decimals",
		"outputs": [
			{
				"internalType": "uint8",
				"name": "",
				"type": "uint8"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "deposit",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "name",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "symbol",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "totalSupply",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "dst",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "wad",
				"type": "uint256"
			}
		],
		"name": "transfer",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "src",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "dst",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "wad",
				"type": "uint256"
			}
		],
		"name": "transferFrom",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "wad",
				"type": "uint256"
			}
		],
		"name": "withdraw",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"stateMutability": "payable",
		"type": "receive"
	}
])
var binDataWbnb ="60c0604052600b60808190527f5772617070656420424e4200000000000000000000000000000000000000000060a090815261003e91600091906100a3565b506040805180820190915260048082527f57424e42000000000000000000000000000000000000000000000000000000006020909201918252610083916001916100a3565b506002805460ff1916601217905534801561009d57600080fd5b5061013e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100e457805160ff1916838001178555610111565b82800160010185558215610111579182015b828111156101115782518255916020019190600101906100f6565b5061011d929150610121565b5090565b61013b91905b8082111561011d5760008155600101610127565b90565b6106728061014d6000396000f3006080604052600436106100ae5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100b8578063095ea7b31461014257806318160ddd1461017a57806323b872dd146101a15780632e1a7d4d146101cb578063313ce567146101e357806370a082311461020e57806395d89b411461022f578063a9059cbb14610244578063d0e30db0146100ae578063dd62ed3e14610268575b6100b661028f565b005b3480156100c457600080fd5b506100cd6102de565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101075781810151838201526020016100ef565b50505050905090810190601f1680156101345780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561014e57600080fd5b50610166600160a060020a036004351660243561036c565b604080519115158252519081900360200190f35b34801561018657600080fd5b5061018f6103d2565b60408051918252519081900360200190f35b3480156101ad57600080fd5b50610166600160a060020a03600435811690602435166044356103d7565b3480156101d757600080fd5b506100b660043561050b565b3480156101ef57600080fd5b506101f86105a0565b6040805160ff9092168252519081900360200190f35b34801561021a57600080fd5b5061018f600160a060020a03600435166105a9565b34801561023b57600080fd5b506100cd6105bb565b34801561025057600080fd5b50610166600160a060020a0360043516602435610615565b34801561027457600080fd5b5061018f600160a060020a0360043581169060243516610629565b33600081815260036020908152604091829020805434908101909155825190815291517fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9281900390910190a2565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103645780601f1061033957610100808354040283529160200191610364565b820191906000526020600020905b81548152906001019060200180831161034757829003601f168201915b505050505081565b336000818152600460209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b303190565b600160a060020a0383166000908152600360205260408120548211156103fc57600080fd5b600160a060020a038416331480159061043a5750600160a060020a038416600090815260046020908152604080832033845290915290205460001914155b1561049a57600160a060020a038416600090815260046020908152604080832033845290915290205482111561046f57600080fd5b600160a060020a03841660009081526004602090815260408083203384529091529020805483900390555b600160a060020a03808516600081815260036020908152604080832080548890039055938716808352918490208054870190558351868152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a35060019392505050565b3360009081526003602052604090205481111561052757600080fd5b33600081815260036020526040808220805485900390555183156108fc0291849190818181858888f19350505050158015610566573d6000803e3d6000fd5b5060408051828152905133917f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65919081900360200190a250565b60025460ff1681565b60036020526000908152604090205481565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103645780601f1061033957610100808354040283529160200191610364565b60006106223384846103d7565b9392505050565b6004602090815260009283526040808420909152908252902054815600a165627a7a723058209ba3a03325af64af3d0633b3b080cc51977daa978f6db1c56ca1dc64f1daa5730029"
//router
var abiDataRouter = JSON.stringify(
  [
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "tokenA",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "tokenB",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "amountADesired",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountBDesired",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountAMin",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountBMin",
          "type": "uint256"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        }
      ],
      "name": "addLiquidity",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "amountA",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountB",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "liquidity",
          "type": "uint256"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "token",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "amountTokenDesired",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountTokenMin",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountETHMin",
          "type": "uint256"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        }
      ],
      "name": "addLiquidityETH",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "amountToken",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountETH",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "liquidity",
          "type": "uint256"
        }
      ],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "tokenA",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "tokenB",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "liquidity",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountAMin",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountBMin",
          "type": "uint256"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        }
      ],
      "name": "removeLiquidity",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "amountA",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountB",
          "type": "uint256"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "token",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "liquidity",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountTokenMin",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountETHMin",
          "type": "uint256"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        }
      ],
      "name": "removeLiquidityETH",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "amountToken",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountETH",
          "type": "uint256"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "token",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "liquidity",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountTokenMin",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountETHMin",
          "type": "uint256"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        }
      ],
      "name": "removeLiquidityETHSupportingFeeOnTransferTokens",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "amountETH",
          "type": "uint256"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "token",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "liquidity",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountTokenMin",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountETHMin",
          "type": "uint256"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        },
        {
          "internalType": "bool",
          "name": "approveMax",
          "type": "bool"
        },
        {
          "internalType": "uint8",
          "name": "v",
          "type": "uint8"
        },
        {
          "internalType": "bytes32",
          "name": "r",
          "type": "bytes32"
        },
        {
          "internalType": "bytes32",
          "name": "s",
          "type": "bytes32"
        }
      ],
      "name": "removeLiquidityETHWithPermit",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "amountToken",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountETH",
          "type": "uint256"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "token",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "liquidity",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountTokenMin",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountETHMin",
          "type": "uint256"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        },
        {
          "internalType": "bool",
          "name": "approveMax",
          "type": "bool"
        },
        {
          "internalType": "uint8",
          "name": "v",
          "type": "uint8"
        },
        {
          "internalType": "bytes32",
          "name": "r",
          "type": "bytes32"
        },
        {
          "internalType": "bytes32",
          "name": "s",
          "type": "bytes32"
        }
      ],
      "name": "removeLiquidityETHWithPermitSupportingFeeOnTransferTokens",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "amountETH",
          "type": "uint256"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "tokenA",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "tokenB",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "liquidity",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountAMin",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountBMin",
          "type": "uint256"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        },
        {
          "internalType": "bool",
          "name": "approveMax",
          "type": "bool"
        },
        {
          "internalType": "uint8",
          "name": "v",
          "type": "uint8"
        },
        {
          "internalType": "bytes32",
          "name": "r",
          "type": "bytes32"
        },
        {
          "internalType": "bytes32",
          "name": "s",
          "type": "bytes32"
        }
      ],
      "name": "removeLiquidityWithPermit",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "amountA",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountB",
          "type": "uint256"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "amountOut",
          "type": "uint256"
        },
        {
          "internalType": "address[]",
          "name": "path",
          "type": "address[]"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        }
      ],
      "name": "swapETHForExactTokens",
      "outputs": [
        {
          "internalType": "uint256[]",
          "name": "amounts",
          "type": "uint256[]"
        }
      ],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "amountOutMin",
          "type": "uint256"
        },
        {
          "internalType": "address[]",
          "name": "path",
          "type": "address[]"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        }
      ],
      "name": "swapExactETHForTokens",
      "outputs": [
        {
          "internalType": "uint256[]",
          "name": "amounts",
          "type": "uint256[]"
        }
      ],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "amountOutMin",
          "type": "uint256"
        },
        {
          "internalType": "address[]",
          "name": "path",
          "type": "address[]"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        }
      ],
      "name": "swapExactETHForTokensSupportingFeeOnTransferTokens",
      "outputs": [],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "amountIn",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountOutMin",
          "type": "uint256"
        },
        {
          "internalType": "address[]",
          "name": "path",
          "type": "address[]"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        }
      ],
      "name": "swapExactTokensForETH",
      "outputs": [
        {
          "internalType": "uint256[]",
          "name": "amounts",
          "type": "uint256[]"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "amountIn",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountOutMin",
          "type": "uint256"
        },
        {
          "internalType": "address[]",
          "name": "path",
          "type": "address[]"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        }
      ],
      "name": "swapExactTokensForETHSupportingFeeOnTransferTokens",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "amountIn",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountOutMin",
          "type": "uint256"
        },
        {
          "internalType": "address[]",
          "name": "path",
          "type": "address[]"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        }
      ],
      "name": "swapExactTokensForTokens",
      "outputs": [
        {
          "internalType": "uint256[]",
          "name": "amounts",
          "type": "uint256[]"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "amountIn",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountOutMin",
          "type": "uint256"
        },
        {
          "internalType": "address[]",
          "name": "path",
          "type": "address[]"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        }
      ],
      "name": "swapExactTokensForTokensSupportingFeeOnTransferTokens",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "amountOut",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountInMax",
          "type": "uint256"
        },
        {
          "internalType": "address[]",
          "name": "path",
          "type": "address[]"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        }
      ],
      "name": "swapTokensForExactETH",
      "outputs": [
        {
          "internalType": "uint256[]",
          "name": "amounts",
          "type": "uint256[]"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "amountOut",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "amountInMax",
          "type": "uint256"
        },
        {
          "internalType": "address[]",
          "name": "path",
          "type": "address[]"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "deadline",
          "type": "uint256"
        }
      ],
      "name": "swapTokensForExactTokens",
      "outputs": [
        {
          "internalType": "uint256[]",
          "name": "amounts",
          "type": "uint256[]"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_factory",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "_WETH",
          "type": "address"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "stateMutability": "payable",
      "type": "receive"
    },
    {
      "inputs": [],
      "name": "factory",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "amountOut",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "reserveIn",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "reserveOut",
          "type": "uint256"
        }
      ],
      "name": "getAmountIn",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "amountIn",
          "type": "uint256"
        }
      ],
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "amountIn",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "reserveIn",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "reserveOut",
          "type": "uint256"
        }
      ],
      "name": "getAmountOut",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "amountOut",
          "type": "uint256"
        }
      ],
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "amountOut",
          "type": "uint256"
        },
        {
          "internalType": "address[]",
          "name": "path",
          "type": "address[]"
        }
      ],
      "name": "getAmountsIn",
      "outputs": [
        {
          "internalType": "uint256[]",
          "name": "amounts",
          "type": "uint256[]"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "amountIn",
          "type": "uint256"
        },
        {
          "internalType": "address[]",
          "name": "path",
          "type": "address[]"
        }
      ],
      "name": "getAmountsOut",
      "outputs": [
        {
          "internalType": "uint256[]",
          "name": "amounts",
          "type": "uint256[]"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "amountA",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "reserveA",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "reserveB",
          "type": "uint256"
        }
      ],
      "name": "quote",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "amountB",
          "type": "uint256"
        }
      ],
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "WETH",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    }
  ]
)
var sc={
  "statusBar":"statusBar",
  "address": "5de5635ede4641b71d466d155a771d7578b086a4",
  "to-address": "",
  "is-deploy": true,
  "amount": "0",
  "name": "cake",
  "abiData": abiDataCake,
  "binData": binDataCake,

}
var messageForm22 = {
  command:"deploy-sc",
  value: sc,
  };
// sendMessage(messageForm22);
//
abiData1={
  
}
var sc1={
  "statusBar":"statusBar",
  "address": "5de5635ede4641b71d466d155a771d7578b086a4",
  "to-address": "",
  "is-deploy": true,
  "amount": "0",
  "name": "wbnb",
  "abiData": abiDataWbnb,
  "binData": binDataWbnb,

}
var messageForm23 = {
  command:"deploy-sc",
  value: sc1,
  };
// sendMessage(messageForm23);


var in11=JSON.stringify({
  "internalType": "address",
  "name": "guy",
  "type": "address",
  "value": "85de3b7a6682c92927098790147f1083ee57ce6f",

});

var in12=JSON.stringify({
  "internalType": "uint256",
  "name": "wad",
  "type": "uint256",
  "value": 9,

});
var in13=JSON.stringify({


  "internalType": "uint256",
  "name": "wad",
  "type": "uint256",
  "value": 9,
});
var in2=JSON.stringify(
  {
    "internalType": "uint256",
    "name": "amountOut",
    "type": "uint256",
    "value": 9,

  },
  {
    "internalType": "address[]",
    "name": "path",
    "type": "address[]",
    "value": 9,

  },
  {
    "internalType": "address",
    "name": "to",
    "type": "address",
    "value": 9,

  },
  {
    "internalType": "uint256",
    "name": "deadline",
    "type": "uint256",
    "value": 9,

  }
);
var balanceOf=JSON.stringify(
  {
    "internalType": "address",
    "name": "account",
    "type": "address",
    "value": "b507d6465c5369459dc5cbc88bf8899e24b7db4c",
  }
)
var mint1=JSON.stringify(
  {
    "internalType": "address",
    "name": "spender",
    "type": "address",
    "value": "b507d6465c5369459dc5cbc88bf8899e24b7db4c",
  },
)
var mint2=JSON.stringify(
  {
    "internalType": "uint256",
    "name": "addedValue",
    "type": "uint256",
    "value": 10000,
  }
)

var approveIn1=JSON.stringify(
  {
    "internalType": "address",
    "name": "spender",
    "type": "address",
    "value": "FF3A857999aBf8c73576e676754983b2F12C76E5",

  },
)
var approveIn2=JSON.stringify(
  {
    "internalType": "uint256",
    "name": "amount",
    "type": "uint256",
    "value": 100000,
  }
)
var allowanceIn1=JSON.stringify(
  {
    "internalType": "address",
    "name": "owner",
    "type": "address",
    "value": "226464747e13b0b120843b18c44d455e868670be",

  },
)
var allowanceIn2=JSON.stringify(
  {
    "internalType": "address",
    "name": "spender",
    "type": "address",
    "value": "FF3A857999aBf8c73576e676754983b2F12C76E5",

  }
)
var functionInputs=[in11,in12]
var functionInputs1=[in13]
var functionInputs2=[in2]
var balanceOfInputs=[balanceOf]
var mintInputs=[mint1,mint2]

var approveInput=[approveIn1,approveIn2]
var allowanceInputs=[allowanceIn1,allowanceIn2]

var encode={
  "function-name":"approve",
  "abiData": abiDataWbnb,
  "inputArray": functionInputs,

}
var encode1={
  "function-name":"withdraw",
  "abiData": abiDataWbnb,
  "inputArray": functionInputs1,

}
var encode2={
  "function-name":"swapETHForExactTokens",
  "abiData": abiDataRouter,
  "inputArray": functionInputs2,

}
var encode3={
  "function-name":"balanceOf",
  "abiData": abiDataCake,
  "inputArray": balanceOfInputs,

}
var encode4={
  "function-name":"allowance",
  "abiData": abiDataCake,
  "inputArray": allowanceInputs,

}


// var messageForm23 = {
//   command:"encode",
//   value: encode4,
//   };
// sendMessage(messageForm23);
//call getOwner of cake
var callGetOwner ={
  'from-address':   "5de5635ede4641b71d466d155a771d7578b086a4",
  'to-address':    "23fb462203cb65eb75b4f3e8e364109ce33e7b8e",
  amount:            "",
  fee:             "",
  tip:             "",
  message:    "",
  'receive-info':  "",
  'is-deploy':   false,
  'is-call':     true,
  name:            "",
  input:           "",
  image:           "",
  "abiData": abiDataCake,
  "function-name":"getOwner",
  isOfflineMode:  false,
  feeType:   false,
  inputArray:"",
  maxGas:"",
  maxGasPriceGwei:"",
  maxTimeUse:"",
  relatedAddresses:[],
}

var messageForm22 = {
  command:"send-transaction",
  value: callGetOwner,
  };
sendMessage(messageForm22);


//call sc mint cake
var callmint ={
  'from-address':   "5de5635ede4641b71d466d155a771d7578b086a4",
  'to-address':    "23fb462203cb65eb75b4f3e8e364109ce33e7b8e",
  amount:            "",
  fee:             "",
  tip:             "",
  message:    "",
  'receive-info':  "",
  'is-deploy':   false,
  'is-call':     true,
  name:            "",
  input:           "",
  image:           "",
  "abiData": abiDataCake,
  "function-name":"mint",
  isOfflineMode:  false,
  feeType:   false,
  inputArray:mintInputs,
  maxGas:"",
  maxGasPriceGwei:"",
  maxTimeUse:"",
  relatedAddresses:[],
}

var messageForm23 = {
  command:"send-transaction",
  value: callmint,
  };
sendMessage(messageForm23);


//call smart contract-get balance -cake

var callsm1 ={
  'from-address':   "5de5635ede4641b71d466d155a771d7578b086a4",
  'to-address':    "23fb462203cb65eb75b4f3e8e364109ce33e7b8e",
  amount:            "",
  fee:             "",
  tip:             "",
  message:    "",
  'receive-info':  "",
  'is-deploy':   false,
  'is-call':     true,
  name:            "",
  input:           "",
  image:           "",
  "abiData": abiDataCake,
  "function-name":"balanceOf",
  isOfflineMode:  false,
  feeType:   false,
  inputArray:balanceOfInputs,
  maxGas:"",
  maxGasPriceGwei:"",
  maxTimeUse:"",
  relatedAddresses:[],
}

var messageForm24 = {
  command:"send-transaction",
  value: callsm1,
  };
sendMessage(messageForm24);

//call smart contract-approve cake

var callsm2 ={
  'from-address':   "e40844ec9eae618baed9c3e5c951e0e022869d97",
  'to-address':    "BFC6af1caCf974d30D65e129654cC5153A56EcF6",
  amount:            "",
  fee:             "1",
  tip:             "",
  message:    "",
  'receive-info':  "",
  'is-deploy':   false,
  'is-call':     true,
  name:            "",
  input:           "",
  image:           "",
  "abiData": abiDataCake,
  "function-name":"approve",
  isOfflineMode:  false,
  feeType:   false,
  inputArray:approveInput,
  maxGas:"",
  maxGasPriceGwei:"",
  maxTimeUse:"",
  relatedAddresses:[],
}

// var messageForm25 = {
//   command:"send-transaction",
//   value: callsm2,
//   };
// sendMessage(messageForm25);

// call smart contract-allowance cake
var callsm3 ={
  'from-address':   "e40844ec9eae618baed9c3e5c951e0e022869d97",
  'to-address':    "BFC6af1caCf974d30D65e129654cC5153A56EcF6",
  amount:            "",
  fee:             "1",
  tip:             "",
  message:    "",
  'receive-info':  "",
  'is-deploy':   false,
  'is-call':     true,
  name:            "",
  input:           "",
  image:           "",
  "abiData": abiDataCake,
  "function-name":"allowance",
  isOfflineMode:  false,
  feeType:   false,
  inputArray:allowanceInputs,
  maxGas:"",
  maxGasPriceGwei:"",
  maxTimeUse:"",
  relatedAddresses:[],
}

// var messageForm26 = {
//   command:"send-transaction",
//   value: callsm3,
//   };
// sendMessage(messageForm26);



// var messageForm6 = {
//   command:"get-all-smart-contract",
// };
// sendMessage(messageForm6);
// //
// var messageForm7 = {
//   command:"get-all-group-d-app",
// };
// sendMessage(messageForm7);
// //
// var messageForm8 = {
//   command:"get-all-d-app-no-group",
// };
// sendMessage(messageForm8);
// //
// // var messageForm9 = {
// //   command:"get-d-app-by-bundle-id",
// //   bundle-id:"com.2048.metaword",
// // };
// // sendMessage(messageForm9);

// //
// var messageForm10 = {
//   command:"get-all-trans",
// };
// sendMessage(messageForm10);
// //
// var messageForm11 = {
//   command:"get-transaction-by-hash",
//   hash:"45b54d4fb5629d0dc50dc5fc3246b97b2a3608c3b70d4948e48d82428df7a82e",
// };
// sendMessage(messageForm11);
// //
// var messageForm12 = {
//   command:"get-transaction-pagination",
//   limit:2,
//   page:1,
// };
// sendMessage(messageForm12);
// //
// var messageForm13 = {
//   command:"get-transaction-by-address-wallet",
//   address:"85de3b7a6682c92927098790147f1083ee57ce6f",
//   limit:10,
//   page:1,
//   status:2,
// };
// sendMessage(messageForm13);
// //
// // connect-node 
// var valueObj2 ={
//   node:                node,
//   wallets:             wallets,
// }

// var messageForm16 = {
//   command:"connect-node",
//   value: valueObj2,
//   };
// sendMessage(messageForm16);
//
// var messageForm17 = {
//   command:"test",
//   address:"1358d7d3bdb6f00c3662ea60b00a75b0541d5c06",
//   };
// sendMessage(messageForm17);
//
// var messageForm18 = {
//   command:"init-connection",
//   address:"1358d7d3bdb6f00c3662ea60b00a75b0541d5c06",
//   };
// sendMessage(messageForm18);

    console.log("sending transaction")
};
