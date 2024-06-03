function uploadFile() {
    const fileInput = document.getElementById('fileInput');
    const file = fileInput.files[0];
    const chunkSize = 1024 * 1024; // 1MB
    const totalChunks = Math.ceil(file.size / chunkSize);
    let uploadedChunks = 0;

    function sendChunk(offset) {
        const xhr = new XMLHttpRequest();
        const chunk = file.slice(offset, offset + chunkSize);
        const formData = new FormData();
        formData.append('chunk', chunk);
        formData.append('fileName', file.name);
        formData.append('offset', offset);
        xhr.open('POST', '/upload', true);
        xhr.onload = function () {
            if (xhr.status == 200) {
                uploadedChunks++;
                //업로드 완료된 경우 페이지 새로고침
                if (uploadedChunks === totalChunks) {
                    alert('Upload complete');
                    location.reload();
                }
            } else {
                alert('Error: ' + xhr.responseText);
            }
        };
        xhr.send(formData);
    }

    for (let offset = 0; offset < file.size; offset += chunkSize) {
        sendChunk(offset);
    }
}

function displayFileName() {
    const input = document.getElementById('fileInput');
    let fileName = '';
    if (input.files.length > 0) {
        fileName = "Selected File: " + input.files[0].name;
    }
    document.getElementById('fileName').textContent = fileName;
}