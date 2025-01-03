require("esbuild").buildSync({
    entryPoints: [
        "./src/online-stats-app.ts",
        "./src/waitlist-stats-app.ts",
        "./src/organizers-companies-app.ts",
        "./src/organizers-vacancies-app.ts",
        "./src/organizers-communities-app.ts",
        "./src/organizers-company-app.ts",
        "./src/companies-and-connections-app.ts",
    ],
    bundle: true,
    minify: process.env.MINIFY === "true",
    outdir: "../public/assets/js",
});
