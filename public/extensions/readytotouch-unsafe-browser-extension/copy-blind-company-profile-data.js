console.log("Blind company profile data copy extension loaded");

// Ctrl + Shift + Y
document.body.addEventListener("keydown", (event) => {
    // Y is for English, Н is for Ukrainian
    if (event.ctrlKey && event.shiftKey && (event.key === "Y" || event.key === "Н")) {
        const goBlindProfileColumns = ``

        navigator.clipboard.writeText(goBlindProfileColumns)
            .then(() => console.log("Page info copied to clipboard:", goBlindProfileColumns))
            .catch((err) => console.error("Failed to copy page info:", err));
    }
});
