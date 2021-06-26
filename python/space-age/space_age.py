class SpaceAge:
    earth_seconds_per_year = 31557600
    planet_orbital_period = {
        "mercury": 0.2408467,
        "venus": 0.61519726,
        "mars": 1.8808158,
        "jupiter": 11.862615,
        "staturn": 29.447498,
        "uranus": 84.016846,
        "neptune": 164.79132,
    }

    def __init__(self, seconds):
        self.seconds = seconds
        self.years = self.seconds / self.earth_seconds_per_year

    def on_earth(self):
        return round(self.years, 2)

    def on_mercury(self):
        return round(self.years / self.planet_orbital_period["mercury"], 2)

    def on_venus(self):
        return round(self.years / self.planet_orbital_period["venus"], 2)

    def on_mars(self):
        return round(self.years / self.planet_orbital_period["mars"], 2)

    def on_jupiter(self):
        return round(self.years / self.planet_orbital_period["jupiter"], 2)

    def on_saturn(self):
        return round(self.years / self.planet_orbital_period["staturn"], 2)

    def on_uranus(self):
        return round(self.years / self.planet_orbital_period["uranus"], 2)

    def on_neptune(self):
        return round(self.years / self.planet_orbital_period["neptune"], 2)
