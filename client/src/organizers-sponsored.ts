export function sponsoredAvailable(optionalPinnedUntil, verified: boolean): boolean {
    if (optionalPinnedUntil === null || !verified) {
        return false;
    }

    const pinnedUntil = new Date(optionalPinnedUntil);
    const today = new Date();
    today.setUTCHours(0, 0, 0, 0);

    return pinnedUntil >= today;
}

export function renderSponsored(entity: string): string {
    return `<div class="card sponsored-card">
    <figure class="sponsored-card__body">
        <img
            alt="Sponsored info icon"
            width="20"
            height="20"
            src="/assets/images/pages/organizer/comment-info.svg"
            class="sponsored-card__image"
        />
        <figcaption class="sponsored-card__caption">
            Want your ${entity} featured at the top? Sponsored placement available — contact us on <a class="link sponsored-card__link" href="https://www.linkedin.com/in/yaroslav-podorvanov/" target="_blank">LinkedIn</a>
        </figcaption>
    </figure>
</div>`;
}
