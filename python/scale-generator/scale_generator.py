class Scale:
    chromatic_sharp = [
        "A",
        "A#",
        "B",
        "C",
        "C#",
        "D",
        "D#",
        "E",
        "F",
        "F#",
        "G",
        "G#",
    ]

    chromatic_flat = ["A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"]

    diatonic_flat = ["F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"]

    def __init__(self, tonic):
        self.tonic = tonic

    def chromatic(self):
        note_in_chromatic_sharp_scale = [
            note for note in Scale.chromatic_sharp if note == self.tonic and note != "C"
        ]
        if note_in_chromatic_sharp_scale:
            return self._generate_chromatic_scale(Scale.chromatic_flat)
        else:
            return self._generate_chromatic_scale(Scale.chromatic_sharp)

    def interval(self, intervals):
        note_in_diatonic_flat_scale = [
            note for note in Scale.diatonic_flat if note == self.tonic
        ]

        if note_in_diatonic_flat_scale:
            return self._generate_interval(Scale.chromatic_flat, intervals)
        return self._generate_interval(Scale.chromatic_sharp, intervals)

    def _generate_chromatic_scale(self, intervals):
        new_scale = []
        i = intervals.index(self.tonic)
        while len(new_scale) < len(intervals):
            new_scale.append(intervals[i % len(intervals)])
            i += 1
        return new_scale

    def _generate_interval(self, scale_type, intervals):

        new_scale = []

        increment_index = {"M": 2, "A": 3, "m": 1}
        i = scale_type.index(self.tonic.capitalize())

        while len(new_scale) < len(intervals):
            for each_step in intervals:
                new_scale.append(scale_type[i % len(scale_type)])
                i += increment_index[each_step]
        return new_scale
