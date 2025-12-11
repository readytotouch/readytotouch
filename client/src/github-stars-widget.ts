function setStarsCount($stars: NodeListOf<Element>, starsCount: number) {
    $stars.forEach(function ($star: Element) {
        $star.innerHTML = starsCount.toString();
        $star.setAttribute("aria-label", `${starsCount} stargazers on GitHub`)
        $star.classList.remove("gh-count--is-hidden");
        $star.removeAttribute("aria-hidden");
    });
}

export function githubStarsWidget() {
    const $stars = document.querySelectorAll(".js-github-count");
    if ($stars.length === 0) {
        return;
    }

    setStarsCount($stars, 1250); // Initial static count

    // @TODO: Uncomment after a successful migration to v3, before starting the article series on Reddit, Hacker News, and forums
    /*
        fetch("https://api.github.com/repos/readytotouch/readytotouch")
            .then(function (response) {
                return response.json();
            })
            .then(function (data) {
                setStarsCount($stars, data.stargazers_count || 0);
            });
        */
}
