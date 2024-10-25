document.addEventListener("DOMContentLoaded", function () {
    chrome.tabs.query({active: true, currentWindow: true}, function (tabs) {
        chrome.scripting.executeScript({
            target: {tabId: tabs[0].id},
            function: checkCompanyStatus
        });
    });

    function checkCompanyStatus() {
        const $status = document.getElementById("js-rtt-extension-popup-status");

        $status.textContent = "popup.js is working!";
    }
});
