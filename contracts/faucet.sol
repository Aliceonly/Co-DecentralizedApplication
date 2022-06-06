// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

contract faucet{
    // address admin=0x00DC6E8B60fA02a5d83e525BBEf3240E8ea54dc5;

    struct users{
         address _address;
         uint  transfercount;//请求的账户获取代币次数
         uint  tokenaccount;//一共获取的币数
    }

    mapping (address=>users) public usertable;
 
    //用户可以通过合约转账1个ether
    function dotransfer(address add)  public{
         require(address(this).balance>=0.01 ether,"no money");
         require(usertable[add].transfercount<3,"today is to many requst,welcome tomorrow");
         uint u = 5 ether;
         payable(add).transfer(u);
         usertable[add]._address=add;
         usertable[add].transfercount++;
         usertable[add].tokenaccount+=5;
  }

//获取账户得币数量
 function gettransfercount(address add) public view returns(uint){
         return usertable[add].transfercount;
 }
 //获取合约余额 
 function getBalancecontract()public view returns(uint){
     return address(this).balance/10**18;
}
//获取本用户余额
 function getBalanceUser()public view returns(uint){
         return msg.sender.balance/10**18;
 }
  fallback() external payable {}
  receive() external payable {}
}