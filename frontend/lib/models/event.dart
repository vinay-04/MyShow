class Event {
  final int id;
  final String name;
  final String date;
  final String location;

  Event(
      {required this.id,
      required this.name,
      required this.date,
      required this.location});

  factory Event.fromJson(Map<String, dynamic> json) {
    return Event(
      id: json['id'],
      name: json['name'],
      date: json['date'],
      location: json['location'],
    );
  }
}
