from locust import HttpUser, task, constant
import os, random

class MyUser(HttpUser):
    wait_time = constant(3)
    host = os.getenv("LOCUST_HOST", "http://waiter-service:8080")

    # List of recipes
    recipes = ['karubikuppa', 'curry', 'spaghetti', 'meuniere', 'sandwich', 'salad', 'smoothie']

    # Random weights for each recipe
    weights = [10, 7, 5, 2, 1, 8, 6]

    @task
    def get_recipe(self):
        headers = {'Content-Type': 'application/json'}
        data = random.choices(self.recipes, weights=self.weights, k=1)[0]
        self.client.get(f"http://waiter-service:8080/{data}", headers=headers)