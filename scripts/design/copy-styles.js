console.log("Copy styles");

const fs = require("fs");

async function extractAndSaveStyles(filesToProcess) {
    const styleRegex = /<style\b[^>]*>([\s\S]*?)<\/style>/gs;

    for (const filePair of filesToProcess) {
        const {source, destination, qtpl_func_name} = filePair;

        try {
            const htmlContent = await fs.promises.readFile(source, "utf8");

            let match;
            const styles = [];

            while ((match = styleRegex.exec(htmlContent)) !== null) {
                styles.push(match[0])
            }

            if (styles.length === 0) {
                console.error(`No styles found in "${source}".`);
                continue;
            }

            const content = [
                `{% func ${qtpl_func_name}() %}`,
                ...styles,
                `{% endfunc %}`,
                "", // Add an empty line at the end
            ]

            await fs.promises.writeFile(destination, content.join("\n"), "utf8");
            console.log(`File "${source}" processed successfully.`);
        } catch (error) {
            console.error(`Error processing file "${source}":`, error);
        }
    }
}

const files = [
    {
        source: './public/design/organizer-go-communities.html',
        destination: './internal/templates/v1/organizers-communities-styles.qtpl',
        qtpl_func_name: 'organizersCommunitiesStyles',
    },
];

async function runScript() {
    await extractAndSaveStyles(files);
}

runScript();
