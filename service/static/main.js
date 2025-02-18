var voteOption = "";
var currentPoll = {};


function loadCaptcha() {
    fetch('/captcha')
        .then(response => response.json())
        .then(data => {
            document.getElementById('captchaImg').src = data.base64Img;
            document.getElementById('captchaID').value = data.id;
        });
}

function loadPool() {
    fetch('/poll')
        .then(response => response.json())
        .then(data => {
            currentPoll = data
            document.getElementById('poll_title').innerText = data.poll.title;

            for (opt of data.poll.options) {
                var div = document.createElement("div");
                div.innerHTML = `<button type="button" onclick="onvote(${opt.index});" name="option" id="opt_${opt.index}" value="${opt.index}">${opt.title}</button>`
                document.getElementById('options').appendChild(div)
            }
        });
}

function loadParcialPool() {
    var total = currentPoll.poll.options.reduce((a, o) => a + (o.quantity || 0), 0)
    document.getElementById('box').style.display = "none"
    document.getElementById('results').style.display = "block"

    for (opt of currentPoll.poll.options) {
        var div = document.createElement("div");
        div.innerHTML = ` <p>${opt.title}  <span>${((opt.quantity || 0)/total * 100).toFixed(2)} % </span></p>`
        document.getElementById('results').appendChild(div)
    }

}

function onvote(option) {
    voteOption = option
    for (opt of currentPoll.poll.options) {
        document.getElementById(`opt_${opt.index}`).className = ""
    }
    document.getElementById(`opt_${option}`).className = "sel"
}

async function sendvote() {
    if(voteOption == "") {
        openModal();
        return
    }

    let captchaID = document.getElementById('captchaID').value;
    let captchaInput = document.getElementById('captchaInput').value;

    const rawResponse = await fetch('/vote', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ captchaID: captchaID, captchaInput: captchaInput, vote: voteOption, pollID: currentPoll.poll.id })
    });
    const content = await rawResponse.json();

    if (content.code == "INVALID_CAPTCHA") {
        document.getElementById('alert').style.display = "block";
        loadCaptcha();
        return
    }

    currentPoll = content
    loadParcialPool();
}

function openModal() {
    document.getElementById('modal').style.display = 'flex';
}

function closeModal() {
    document.getElementById('modal').style.display = 'none';
}

window.onload = function () {
    loadCaptcha();
    loadPool();

    votingForm.onkeypress = function (key) {
        let btn = 0 || key.keyCode || key.charCode;
        if (btn == 13) {
            key.preventDefault();
            return false;
        }

    }
}        