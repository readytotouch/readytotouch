const overOneMonth = "Over one month";
const overTwoMonths = "Over two months";
const overThreeMonths = "Over three months";
const overFourMonths = "Over four months";
const overFiveMonths = "Over five months";
const overSixMonths = "Over six months";
const overSevenMonths = "Over seven months";
const overEightMonths = "Over eight months";
const overNineMonths = "Over nine months";
const overTenMonths = "Over ten months";
const overElevenMonths = "Over eleven months";
const overAYear = "Over a year";

class VacancyPeriod {
    constructor(
        public readonly name: string,
        public readonly before: Date,
        public used: boolean = false,
    ) {
    }
}

export class VacancyPeriodContainer {
    private periods: Array<VacancyPeriod>;

    constructor() {
        const today = truncateToUtcDay(new Date());

        this.periods = [
            new VacancyPeriod(overOneMonth, addMonths(today, -1)),
            new VacancyPeriod(overTwoMonths, addMonths(today, -2)),
            new VacancyPeriod(overThreeMonths, addMonths(today, -3)),
            new VacancyPeriod(overFourMonths, addMonths(today, -4)),
            new VacancyPeriod(overFiveMonths, addMonths(today, -5)),
            new VacancyPeriod(overSixMonths, addMonths(today, -6)),
            new VacancyPeriod(overSevenMonths, addMonths(today, -7)),
            new VacancyPeriod(overEightMonths, addMonths(today, -8)),
            new VacancyPeriod(overNineMonths, addMonths(today, -9)),
            new VacancyPeriod(overTenMonths, addMonths(today, -10)),
            new VacancyPeriod(overElevenMonths, addMonths(today, -11)),
            new VacancyPeriod(overAYear, addMonths(today, -12)),
        ];
    }

    public over(vacancyDate: Date): string[] {
        const names: string[] = [];

        for (const p of this.periods) {
            if (!p.used && vacancyDate < p.before) {
                names.push(p.name);
                p.used = true;
            }
        }

        return names;
    }
}

function truncateToUtcDay(date: Date): Date {
    const d = new Date(date);
    d.setUTCHours(0, 0, 0, 0);
    return d;
}

function addMonths(date: Date, months: number): Date {
    const d = new Date(date);
    d.setUTCMonth(d.getUTCMonth() + months);
    return d;
}
