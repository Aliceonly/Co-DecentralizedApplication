if (typeof window.ethereum !== 'undefined') {
    console.log('MetaMask is installed!')
      getAccount();
      ethereum.on('accountsChanged', function (accounts) {
          // Time to reload your interface with accounts[0]!
          console.log("现在的账户是:", accounts[0]);
          account=accounts[0];
          $("#address").html(accounts[0]);
      });
  } else {
    alert('plase install the MetaMask')
  }
  var account;
  async function getAccount() {
      yghaccounts = await ethereum.request({ method: 'eth_requestAccounts' });
      account = yghaccounts[0];
     $("#address").html(account);
     console.log(account)
  }