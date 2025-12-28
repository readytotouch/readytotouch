export function parseCurrentProgrammingLanguage(pathname: string): string {
    const index = pathname.indexOf("/", 2);

    return pathname.substring(0, index);
}
