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
    {
        source: './public/design/organizer-vacancies-subscribe.html',
        destination: './internal/templates/v1/organizers-waitlist-styles.qtpl',
        qtpl_func_name: 'organizersWaitlistStyles',
    },
    {
        source: './public/design/organizer-companies.html',
        destination: './internal/templates/v1/organizers-companies-v2-styles.qtpl',
        qtpl_func_name: 'organizersCompaniesV2Styles',
    },
    {
        source: './public/design/organizer-statistics.html',
        destination: './internal/templates/v1/organizers-company-v2-styles.qtpl',
        qtpl_func_name: 'organizersCompanyV2Styles',
    },
    {
        source: './public/design/organizer-vacancies.html',
        destination: './internal/templates/v1/organizers-vacancies-v2-styles.qtpl',
        qtpl_func_name: 'organizersVacanciesV2Styles',
    },
    {
        source: './public/design/organizer-welcome.html',
        destination: './internal/templates/v1/organizers-welcome-styles.qtpl',
        qtpl_func_name: 'organizersWelcomeStyles',
    },
    {
        source: './public/design/organizer-main-page.html',
        destination: './internal/templates/v1/organizers-main-styles.qtpl',
        qtpl_func_name: 'organizersMainStyles',
    },
    {
        source: './public/design/online-new.html',
        destination: './internal/templates/v1/organizers-online-styles.qtpl',
        qtpl_func_name: 'organizersOnlineStyles',
    },
    {
        source: './public/design/connections.html',
        destination: './internal/templates/v1/cac-companies-and-connections-v1-styles.qtpl',
        qtpl_func_name: 'companiesAndConnectionsV1Styles',
    },
    /*
        {
            source: './public/design/',
            destination: './internal/templates/v1/',
            qtpl_func_name: '',
        },
    */
];

async function runScript() {
    await extractAndSaveStyles(files);
}

runScript();
