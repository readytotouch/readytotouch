console.log("LinkedIn company profile data copy extension loaded");

// Ctrl + Shift + Y
document.body.addEventListener("keydown", (event) => {
    if (event.ctrlKey && event.shiftKey && event.key === "Y") {
        const pageInfo = {
            url: window.location.href,
            name: document.querySelector("h1")?.innerText || "",
        };

        const jsonString = JSON.stringify(pageInfo);

        navigator.clipboard.writeText(jsonString)
            .then(() => console.log("Page info copied to clipboard:", jsonString))
            .catch((err) => console.error("Failed to copy page info:", err));
    }
});
