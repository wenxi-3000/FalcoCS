async function GenerateBinary() {
    let address = document.getElementById('address');
    let port = document.getElementById('port');
    let filename = document.getElementById('filename');

    if (!address.value || !port.value || !filename.value) {
        ShowNotification('warning', 'Ops!', 'You should fill all the required fields.');
        return
    }

    Swal.fire({
        title: 'Building...',
        onBeforeOpen: () => {
            Swal.showLoading()
        }
    });

    generate(address.value, port.value, filename.value)
}

async function generate(address, port, filename) {
    var httpRequest = new XMLHttpRequest();	//第一步：创建需要的对象
    const url = '/generate';
    httpRequest.open('POST', url, true);
    httpRequest.setRequestHeader("Content-type","application/json; charset=utf-8");	// 设置请求头,注：post方式必须设置请求头（在建立连接后设置请求头）
    var obj = {"address": address, "port": port, "filename": filename};
    httpRequest.send(JSON.stringify(obj));//发送请求 将json写入send中
    httpRequest.onreadystatechange = function () {	// 第六步：处理接收到的数据 请求后的回调接口，可将请求成功后要执行的程序写在其中
        if (httpRequest.readyState == 4 && httpRequest.status == 200) {	// 第七步：验证请求是否发送成功
            var recvText = httpRequest.responseText;	// 第八步：获取到服务端返回的数据
            window.location.href = "download/" + recvText;
        }
    };

}