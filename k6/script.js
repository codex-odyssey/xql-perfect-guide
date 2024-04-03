import http from 'k6/http';

const baseEndpoint = __ENV.BASE_URL || "http://localhost:8080"
const recipes = [
    "karubikuppa",
    "curry",
    "spaghetti",
    "meuniere",
    "sandwich",
    "salad",
    "smoothie"
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
    const recipe = recipes[Math.floor(Math.random() * recipes.length)]
    http.get(`${baseEndpoint}/${recipe}`);
}
