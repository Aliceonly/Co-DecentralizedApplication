// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;
import "github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v3.3/contracts/cryptography/ECDSA.sol";

contract TaskDeployer {
     using ECDSA for bytes32;
     //创建任务的事件
     event Creat_Event(address add,string taskname,string taskcategory ,uint amount,uint timestamp);
     //接受任务的事件
     event Confirm_Event(address add,uint timestamp);
     //取消任务的事件
     event  Cancel_Event(address add,uint timestamp);
     //确认任务完成事件
     event  claim_Event(address add,uint timestamp);
     //创建用户事件
     event Creatuser_Event(string name,string phonenumber,uint studentid);
     //任务的结构体
     struct Task{
        string Taskname;//任务的名称
        address sponsor;//任务的发起者
        address beneficiary;//接受任务的人
        string  category;//任务的分类  
        uint   amount;//任务的报酬
        uint timestamp;//任务的时间戳
        string state;//任务的状态
        string launchTime;//任务的工作时间
     }
     //用户结构体
     struct User{
         string name;
         string phonenumber;
         uint   studentid;
         string password;

     }
     //所有任务列表
     mapping(uint => Task) public _tasklist; //使用时间戳对应任务
     mapping(uint => User)public _userlist;//使用学生证号对应
     /**
     * @dev 检索调用者是否为具体任务的受益人
     */
    modifier taskBelongsTo(uint _timestamp) {
        bool found = false;
        if(_tasklist[_timestamp].beneficiary==msg.sender){
            found=true;
        }
        require(found == true, "You don't have permission to do so.");
        _;
    }

    /**
     * @dev 检索任务是否为发布者
     */
    modifier issponsoradress(uint _timestamp) {
        bool found = false;
       if(_tasklist[_timestamp].sponsor==msg.sender){
           found=true;
       }
        require(found == false, "The beneficiary must be unique.");
        _;
    }

    //比较字符串相等函数
    function stringsEquals(string memory s1, string memory s2) private pure returns (bool) {
    bytes memory b1 = bytes(s1);
    bytes memory b2 = bytes(s2);
    uint256 l1 = b1.length;
    if (l1 != b2.length) return false;
    for (uint256 i=0; i<l1; i++) {
        if (b1[i] != b2[i]) return false;
    }
    return true;
}

 //实现用户提前打钱到合约在通过合约打钱到用户的操作
            // 向合约账户转账 
            function transderToContract() payable public {
                payable(address(this)).transfer(msg.value);
            }
            
            // 获取合约账户余额 
            function getBalanceOfContract() public view returns (uint256) {
                return address(this).balance;
            }
            //实现合约转账给完成任务的用户或者发布任务者取消任务金额退返操作
             function svalue(address payable addr,uint amount)public payable{      
                addr.transfer(amount* 10**18);
    }
            

     /*
     创建任务
     */
      uint _now;
      function createNewEvent(string memory _launchTime, string memory _category,string memory _taskname,uint _amount) public payable returns(uint) {
        _now = block.timestamp;
        Task storage t = _tasklist[_now];
        t.Taskname=_taskname;
        t.sponsor=msg.sender;//任务的发起者
        t.beneficiary=address(0);//接受任务的人
        t.category=_category;//任务的分类  
        t.amount=_amount;//任务的报酬
        t.timestamp=_now;//任务的时间戳
        t.state="availablev";//任务的状态
        t.launchTime=_launchTime;//任务的工作时间
        emit Creat_Event(msg.sender,_taskname,_category,_amount,_now);
        transderToContract();
        return _now;
    }
     function query()public view returns(uint){
        return _now;
     }  
   /*
     接受任务
     */
    function confirmtask(uint _timestamp) public  {
        require(_tasklist[_timestamp].beneficiary==address(0),"this task was confirm");
        require(_tasklist[_timestamp].sponsor!=msg.sender,"The publisher cannot accept it");
        require(stringsEquals(_tasklist[_timestamp].state,"invalid")!=true,"this task is cancel");
        if(stringsEquals(_tasklist[_timestamp].state,"invalid") ||stringsEquals(_tasklist[_timestamp].state,"completed") !=true){
            _tasklist[_timestamp].beneficiary=msg.sender;
           emit Confirm_Event(msg.sender,block.timestamp);
        }else{
            revert("the task have error");
        }
    }
     /*
     取消任务
     */
    function cancelEvent(uint _timestamp) public {
       require(_tasklist[_timestamp].sponsor == msg.sender,"must be sponsor can cancel");
        uint amount = _tasklist[_timestamp].amount;
        // uint256 contractbalance=getBalanceOfContract();
        // require(contractbalance==amount,"the contractbanlance have error");
        _tasklist[_timestamp].state="invalid";
        svalue(payable(msg.sender),amount);
        emit Cancel_Event(msg.sender,block.timestamp);
    }

      /*
    发布任务者确认接受任务者完成任务
     */
      function claimTrust(uint _timestamp,bytes memory _sigs,bytes32 _PHash,string memory _taskname) public issponsoradress(_timestamp) {
        require(stringsEquals(_tasklist[_timestamp].state,"invalid")!=true, "Task is invalid.");
        require(stringsEquals(_tasklist[_timestamp].state,"completed")!=true, "Task is completed.");
        require(_checkSigs(_timestamp, _sigs, _PHash,_taskname), "invalid sig");

        uint amount = _tasklist[_timestamp].amount;

        // 更新任务状态
        _tasklist[_timestamp].state ="completed";
        _tasklist[_timestamp].amount = 0;

        // 将交易金额发送给受益人
        address payable beneficiary = payable(_tasklist[_timestamp].beneficiary);
              svalue(beneficiary,amount);
        emit Confirm_Event(msg.sender,block.timestamp);
    }

     /*
     存储用户信息
     */
     function createuser(string memory name,string memory phonenumber,uint studentid,string memory password) public {
          User storage u = _userlist[studentid];
          u.name=name;
          u.phonenumber=phonenumber;
          u.studentid=studentid;
          u.password=password;
          emit Creatuser_Event(name,phonenumber,studentid);
     }
      /*
     用户修改信息
     */
     //用户修改名称
     function changename(string memory _name,uint _studentid,string memory _password) public {
         require(
            stringsEquals(_userlist[_studentid].password,_password)==true,"please sesure your password"
         );
         _userlist[_studentid].name=_name;
     }
     //用户修改密码
     function changepassword(string memory _password,uint _studentid) public{
         require(stringsEquals(_userlist[_studentid].password,_password)!=true,"password is simply");
         _userlist[_studentid].password=_password;
     }

  //验证签名发布任务者确认接收者完成任务
      //获取任务的哈希
     function getTxHash(
        string memory _taskname,
        string memory _timestamp
    ) public view returns (bytes32) {
        return keccak256(abi.encodePacked(address(this), _taskname, _timestamp));
    }
     //验证签名的正确性
    function _checkSigs(uint _timestamp,bytes memory _sigs, bytes32 _txHash,string memory _taskname)
        private
        view
        returns (bool)
    {
        if(stringsEquals(_tasklist[_timestamp].Taskname,_taskname)==true){

        bytes32 ethSignedHash = _txHash.toEthSignedMessageHash();
        for (uint i = 0; i < _sigs.length; i++) {
            address signer = ethSignedHash.recover(_sigs);
            bool valid = signer == msg.sender;
            if (!valid) {
                return false;
            }
        }
        return true;
        }else{
          revert("false");
        }
    }

   fallback() external payable {}
   receive() external payable {}
     

}