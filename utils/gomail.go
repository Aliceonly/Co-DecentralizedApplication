package utils
 
import (
    "fmt"
    "log"
    "math/rand"
    "time"
 
    "gopkg.in/gomail.v2"
)
 
// MailboxConf 邮箱配置
type MailboxConf struct {
    // 邮件标题
    Title string
    // 邮件内容
    Body string
    // 收件人列表
    RecipientList []string
    // 发件人账号
    Sender string
    // 发件人密码，授权码
    SPassword string
    // SMTP 服务器地址
    SMTPAddr string
    // SMTP端口
    SMTPPort int
}
 
func SendVcode() (vcode string) {
    var mailConf MailboxConf
    mailConf.Title = "来自 Dapp项目组 的一封信"
  
    // mailConf.RecipientList = []string{"邮箱账号1","邮箱账号2"}
    mailConf.Sender = `aliceonly@foxmail.com`
  
  //填写授权码，邮箱密码
    mailConf.SPassword = "cqictojijjuagiag"
  
    // QQ邮箱：SMTP服务器地址：smtp.qq.com（端口：587）
    // 雅虎邮箱: SMTP服务器地址：smtp.yahoo.com（端口：587）
    // 163邮箱：SMTP服务器地址：smtp.163.com（端口：25）
    // 126邮箱: SMTP服务器地址：smtp.126.com（端口：25）
    // 新浪邮箱: SMTP服务器地址：smtp.sina.com（端口：25）
 
    mailConf.SMTPAddr = `smtp.qq.com`
    mailConf.SMTPPort = 587
 
    //产生六位数验证码
    rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
    vcode = fmt.Sprintf("%06v", rnd.Int31n(1000000))
 
    //发送的内容
    html := fmt.Sprintf(`<div>
        <div>
            尊敬的用户%s，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>您本次的邮箱验证码为%s , 为了保证您的账号安全 , 验证码有效期为5分钟。请确认为本人操作 , 切勿向他人泄露 , 感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮件由Dapp项目组邮箱验证系统自动发出 , 请勿回复。</p>
        </div>
    </div>`, mailConf.Sender, vcode)
 
    m := gomail.NewMessage()
  
    m.SetHeader(`From`,mailConf.Sender, "Dapp项目组")
    m.SetHeader(`To`, mailConf.Sender)
    m.SetHeader(`Subject`, mailConf.Title)
    m.SetBody(`text/html`, html)
    // m.Attach("./Dockerfile") //添加附件
    err := gomail.NewDialer(mailConf.SMTPAddr, mailConf.SMTPPort, mailConf.Sender, mailConf.SPassword).DialAndSend(m)
    if err != nil {
        log.Fatalf("邮件递送失败, %s", err.Error())
        return
    }
    log.Printf("验证邮件已发出")

    return
}