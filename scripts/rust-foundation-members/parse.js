/**
 * https://rustfoundation.org/members/
 */

{
    const companies = Array.from(document.querySelectorAll(".cmp-modal__header")).map(function ($header) {
        return {
            "name": $header.querySelector("h3").textContent.trim(),
            "website": $header.querySelector("a")?.href,
        };
    }).filter(_ => _.website);

    function generateMarkdownTable(companies) {
        const markdown = [
            `| Name | Website | LinkedIn Name | LinkedIn Website |`,
            `|------|---------|--------------|-----------------|`,
        ];

        companies.forEach(company => {
            const linkedInName = `[LinkedIn ${company.name}](https://www.google.com/search?q=site%3Alinkedin.com+${encodeURIComponent(company.name)})`;
            const linkedInWebsite = `[LinkedIn ${company.website}](https://www.google.com/search?q=site%3Alinkedin.com+${encodeURIComponent(company.website)})`;

            markdown.push(`| ${company.name} | [${company.website}](${company.website}) | ${linkedInName} | ${linkedInWebsite} |`);
        });

        return markdown.join("\n");
    }

    console.log(generateMarkdownTable(companies));
}
