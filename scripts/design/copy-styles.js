console.log("Copy styles");

const fs = require('fs');

/**
 * Парсить вихідний файл на наявність DOM-елементів <style>
 * і зберігає їх у файл призначення разом з тегом <style>,
 * використовуючи лише вбудовані модулі Node.js (регулярні вирази).
 *
 * УВАГА: Цей підхід є КРИХКИМ і може працювати неправильно
 * для складного або погано сформованого HTML.
 * Регулярні вирази не є надійним способом парсингу HTML.
 * Для надійного парсингу HTML рекомендовано використовувати бібліотеки,
 * такі як jsdom або cheerio.
 *
 * @param {Array<Object>} filesToProcess - Масив об'єктів { source, destination }.
 */
async function extractAndSaveStylesNative(filesToProcess) {
    // Регулярний вираз для пошуку тегів <style> та їхнього вмісту.
    // 's' флаг дозволяє '.' співставляти символи нового рядка.
    // 'g' флаг дозволяє знайти всі співпадіння, а не лише перше.
    const styleRegex = /<style\b[^>]*>([\s\S]*?)<\/style>/gs;

    for (const filePair of filesToProcess) {
        const { source, destination } = filePair;

        try {
            const htmlContent = await fs.promises.readFile(source, 'utf8');

            let match;
            let stylesToSave = '';
            let foundStyles = false;

            // Використання exec() в циклі для пошуку всіх співпадінь
            while ((match = styleRegex.exec(htmlContent)) !== null) {
                // match[0] містить весь знайдений тег <style> з вмістом
                // match[1] містить лише вміст (textContent) всередині тегу <style>
                stylesToSave += match[0] + '\n'; // Зберігаємо весь тег <style>
                foundStyles = true;
            }

            if (!foundStyles) {
                console.log(`У файлі "${source}" не знайдено елементів <style> (нативним методом).`);
                continue;
            }

            await fs.promises.writeFile(destination, stylesToSave.trim(), 'utf8');
            console.log(`Стилі з "${source}" успішно збережено у "${destination}" (нативним методом).`);

        } catch (error) {
            console.error(`Помилка обробки файлу "${source}" (нативним методом):`, error.message);
        }
    }
}

// --- Приклад використання ---
const myFiles = [
    { source: './index.html', destination: './extracted-styles-native.css' },
    { source: './about.html', destination: './about-styles-native.css' },
];

// Створюємо тестові файли для демонстрації (можна пропустити, якщо вже є)
async function createTestFiles() {
    await fs.promises.writeFile('./index.html', `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Test Page 1</title>
            <style id="main-style">
                body {
                    font-family: Arial, sans-serif;
                    margin: 20px;
                }
                h1 {
                    color: blue;
                }
            </style>
        </head>
        <body>
            <h1>Hello from Index!</h1>
            <p>Some content here.</p>
            <style>
                p {
                    font-size: 16px;
                }
            </style>
        </body>
        </html>
    `, 'utf8');

    await fs.promises.writeFile('./about.html', `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Test Page 2</title>
            <style>
                .container {
                    padding: 10px;
                    border: 1px solid green;
                }
            </style>
            <style type="text/css">
                /* Ще один стиль */
                .another-class {
                    background-color: lightgray;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <h2>About Us</h2>
            </div>
        </body>
        </html>
    `, 'utf8');
    console.log('Тестові файли index.html та about.html створено для нативного парсингу.');
}


// Запуск скрипта
async function runNativeScript() {
    await createTestFiles(); // Закоментуйте, якщо не потрібно створювати тестові файли
    await extractAndSaveStylesNative(myFiles);
}

runNativeScript();
