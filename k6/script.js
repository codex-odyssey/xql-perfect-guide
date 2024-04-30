import { sleep } from 'k6';
import http from 'k6/http';

const baseEndpoint = __ENV.BASE_URL || 'http://localhost:8080'

// List of recipes and weights
const items = [
    { name: 'karubikuppa', weight: 20 },
    { name: 'curry', weight: 7 },
    { name: 'spaghetti', weight: 5 },
    { name: 'meuniere', weight: 2 },
    { name: 'sandwich', weight: 1 },
    { name: 'salad', weight: 15 },
    { name: 'smoothie', weight: 6 },
    { name: 'yakitori', weight: 8 },
    { name: 'yasaiitame', weight: 3 },
    { name: 'yakiniku', weight: 1 }
]

export const options = {
    stages: [
        { duration: '3m', target: 10 },
        { duration: '10m', target: 20 },
        { duration: '1m', target: 0 },
    ],
    noConnectionReuse: true,
};

export default function () {
    const recipe = getRandomItemWeighted(items);
    sleep(1)
    http.get(`${baseEndpoint}/${recipe}`);
}

function getRandomItemWeighted(items) {
    const totalWeight = items.reduce((acc, item) => acc + item.weight, 0);
    const randomWeight = Math.random() * totalWeight;
    let cumulativeWeight = 0;
    for (const item of items) {
        cumulativeWeight += item.weight;
        if (randomWeight <= cumulativeWeight) {
            return item.name;
        }
    }
}
