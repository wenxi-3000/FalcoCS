async function GenerateBinary() {
    let address = document.getElementById('address');
    let port = document.getElementById('port');
    let filename = document.getElementById('filename');

    if (!address.value || !port.value) {
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
        .then(response => {
            if (!response.ok) {
                return response.text().then(err => {
                    throw new Error(err);
                });
            }
            return response.text();
        })
        .then(response => {
            Swal.close();
            window.location.href = 'download/' + response;
        })
        .catch(err => {
            console.log('Error: ', err);
            Swal.close();
            ShowNotification('danger', 'Ops!', 'Failed building client binary.\n' + JSON.parse(err.message).error)
        });
}

async function generate(address, port, filename) {
    event.preventDefault();
    let formData = new FormData();
    formData.append('address', address);
    formData.append('port', port);
    formData.append('filename', filename);

    const url = '/generate';
    const initDetails = {
        method: 'POST',
        body: formData,
        mode: "cors",
    }

    let response = await fetch(url, initDetails);
    let data = await response;
    return data;
}