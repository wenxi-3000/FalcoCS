// Add event when `enter` key was pressed on username
window.addEventListener("load", function () {
    let input = document.getElementById("inputUsername");
    input.addEventListener("keyup", function (event) {
        if (event.keyCode === 13) {
            event.preventDefault();
            document.getElementById("loginButton").click();
        }
    });
});

// Add event when `enter` key was pressed on password
window.addEventListener("load", function () {
    let input = document.getElementById("inputPassword");
    input.addEventListener("keyup", function (event) {
        if (event.keyCode === 13) {
            event.preventDefault();
            document.getElementById("loginButton").click();
        }
    });
});

function Login() {
    let username = document.getElementById('inputUsername');
    let password = document.getElementById('inputPassword');
    if (!username.value || !password.value) {
        return
    }

    auth(username.value, password.value)
}

async function auth(username, password) {
    // let formData = new FormData();

    // formData.append('username', username);
    // formData.append('password', password);
    var httpRequest = new XMLHttpRequest();	//第一步：创建需要的对象
    const url = '/auth';
    httpRequest.open('POST', url, true);
    httpRequest.setRequestHeader("Content-type","application/json; charset=utf-8");	// 设置请求头,注：post方式必须设置请求头（在建立连接后设置请求头）
    var obj = {"username": username, "password": password};
    httpRequest.send(JSON.stringify(obj));//发送请求 将json写入send中
    httpRequest.onreadystatechange = function () {	// 第六步：处理接收到的数据 请求后的回调接口，可将请求成功后要执行的程序写在其中
        if (httpRequest.readyState == 4 && httpRequest.status == 200) {	// 第七步：验证请求是否发送成功
            var recvText = httpRequest.responseText;	// 第八步：获取到服务端返回的数据
            var recvJson = JSON.parse(recvText);	// 将 json字符串 转化为 json对象
            console.log(recvText);
            console.log(recvJson['data']['x_token']);
            window.location.href = "/";
        }
    };
}