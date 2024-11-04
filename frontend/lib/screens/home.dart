import 'package:flutter/material.dart';
import 'package:myshow/components/eventCard.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  _HomeScreenState createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      floatingActionButton: FloatingActionButton.extended(
        onPressed: () {
          showModalBottomSheet(
            showDragHandle: true,
            context: context,
            builder: (BuildContext context) {
              return SingleChildScrollView(
                child: Container(
                  padding: const EdgeInsets.all(24.0),
                  child: Column(
                    mainAxisSize: MainAxisSize.min,
                    crossAxisAlignment: CrossAxisAlignment.stretch,
                    children: [
                      Text(
                        'Add New Event',
                        style: TextStyle(
                          fontSize: 24,
                          fontWeight: FontWeight.bold,
                          color: Colors.purple[700],
                        ),
                      ),
                      const SizedBox(height: 24),
                      TextField(
                        decoration: InputDecoration(
                          labelText: 'Title',
                          labelStyle: TextStyle(color: Colors.purple[700]),
                          enabledBorder: OutlineInputBorder(
                            borderSide: BorderSide(color: Colors.grey[300]!),
                            borderRadius: BorderRadius.circular(8),
                          ),
                          focusedBorder: OutlineInputBorder(
                            borderSide: BorderSide(color: Colors.purple[700]!),
                            borderRadius: BorderRadius.circular(8),
                          ),
                        ),
                      ),
                      const SizedBox(height: 16),
                      TextField(
                        decoration: InputDecoration(
                          labelText: 'Location',
                          labelStyle: TextStyle(color: Colors.purple[700]),
                          enabledBorder: OutlineInputBorder(
                            borderSide: BorderSide(color: Colors.grey[300]!),
                            borderRadius: BorderRadius.circular(8),
                          ),
                          focusedBorder: OutlineInputBorder(
                            borderSide: BorderSide(color: Colors.purple[700]!),
                            borderRadius: BorderRadius.circular(8),
                          ),
                          prefixIcon: const Icon(Icons.location_on_outlined),
                        ),
                      ),
                      const SizedBox(height: 16),
                      TextField(
                        decoration: InputDecoration(
                          labelText: 'Date',
                          labelStyle: TextStyle(color: Colors.purple[700]),
                          enabledBorder: OutlineInputBorder(
                            borderSide: BorderSide(color: Colors.grey[300]!),
                            borderRadius: BorderRadius.circular(8),
                          ),
                          focusedBorder: OutlineInputBorder(
                            borderSide: BorderSide(color: Colors.purple[700]!),
                            borderRadius: BorderRadius.circular(8),
                          ),
                          prefixIcon: const Icon(Icons.calendar_today),
                        ),
                        readOnly: true,
                        onTap: () async {
                          DateTime? pickedDate = await showDatePicker(
                            context: context,
                            initialDate: DateTime.now(),
                            firstDate: DateTime(2000),
                            lastDate: DateTime(2101),
                            builder: (context, child) {
                              return Theme(
                                data: Theme.of(context).copyWith(
                                  colorScheme: ColorScheme.light(
                                    primary: Colors.purple[700]!,
                                  ),
                                ),
                                child: child!,
                              );
                            },
                          );
                          if (pickedDate != null) {
                            // Handle the picked date
                          }
                        },
                      ),
                      const SizedBox(height: 16),
                      TextField(
                        decoration: InputDecoration(
                          labelText: 'Time',
                          labelStyle: TextStyle(color: Colors.purple[700]),
                          enabledBorder: OutlineInputBorder(
                            borderSide: BorderSide(color: Colors.grey[300]!),
                            borderRadius: BorderRadius.circular(8),
                          ),
                          focusedBorder: OutlineInputBorder(
                            borderSide: BorderSide(color: Colors.purple[700]!),
                            borderRadius: BorderRadius.circular(8),
                          ),
                          prefixIcon: const Icon(Icons.access_time),
                        ),
                        readOnly: true,
                        onTap: () async {
                          TimeOfDay? pickedTime = await showTimePicker(
                            context: context,
                            initialTime: TimeOfDay.now(),
                            builder: (context, child) {
                              return Theme(
                                data: Theme.of(context).copyWith(
                                  colorScheme: ColorScheme.light(
                                    primary: Colors.purple[700]!,
                                  ),
                                ),
                                child: child!,
                              );
                            },
                          );
                          if (pickedTime != null) {
                            // Handle the picked time
                          }
                        },
                      ),
                      const SizedBox(height: 16),
                      TextField(
                        maxLines: 3,
                        decoration: InputDecoration(
                          labelText: 'Description',
                          labelStyle: TextStyle(color: Colors.purple[700]),
                          enabledBorder: OutlineInputBorder(
                            borderSide: BorderSide(color: Colors.grey[300]!),
                            borderRadius: BorderRadius.circular(8),
                          ),
                          focusedBorder: OutlineInputBorder(
                            borderSide: BorderSide(color: Colors.purple[700]!),
                            borderRadius: BorderRadius.circular(8),
                          ),
                        ),
                      ),
                      const SizedBox(height: 32),
                      ElevatedButton(
                        style: ElevatedButton.styleFrom(
                          backgroundColor: Colors.purple[700],
                          padding: const EdgeInsets.symmetric(vertical: 16),
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(8),
                          ),
                        ),
                        onPressed: () {
                          // Handle the save event action
                        },
                        child: const Text(
                          'Save Event',
                          style: TextStyle(
                            fontSize: 16,
                            fontWeight: FontWeight.bold,
                            color: Colors.white,
                          ),
                        ),
                      ),
                    ],
                  ),
                ),
              );
            },
          );
        },
        label: const Text('Add Event',
            style: TextStyle(color: Colors.white, fontSize: 24)),
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(4),
        ),
        backgroundColor: Colors.purple[700],
      ),
      floatingActionButtonLocation: FloatingActionButtonLocation.centerFloat,
      backgroundColor: Colors.grey[100],
      appBar: AppBar(
        title: const Text(
          'My Show',
          style: TextStyle(color: Colors.white, fontWeight: FontWeight.bold),
        ),
        backgroundColor: Colors.purple[700],
        elevation: 0,
        actions: [
          IconButton(
            icon: const Icon(
              Icons.person_outlined,
              color: Colors.white,
              size: 28,
            ),
            onPressed: () {
              Navigator.pushNamed(context, '/profile');
            },
          ),
        ],
      ),
      body: CustomScrollView(
        slivers: [
          SliverToBoxAdapter(
            child: Padding(
              padding: const EdgeInsets.all(16.0),
              child: Row(
                children: [
                  Text(
                    'Upcoming Events',
                    style: TextStyle(
                      fontSize: 24,
                      fontWeight: FontWeight.bold,
                      color: Colors.purple[800],
                    ),
                  ),
                  const Spacer(),
                  IconButton(
                    icon: const Icon(Icons.sort),
                    onPressed: () {
                      final RenderBox sortByBox =
                          context.findRenderObject() as RenderBox;
                      final Offset sortbyPos =
                          sortByBox.localToGlobal(Offset.zero);
                      final RenderBox overlay = Overlay.of(context)
                          .context
                          .findRenderObject() as RenderBox;

                      showMenu(
                          context: context,
                          position: RelativeRect.fromRect(
                            Rect.fromLTWH(
                              sortbyPos.dx + sortByBox.size.width - 56,
                              sortbyPos.dy + 120,
                              0,
                              0,
                            ),
                            Offset.zero & overlay.size,
                          ),
                          items: [
                            const PopupMenuItem(
                              value: 'date',
                              child: Text('Sort by Date'),
                            ),
                            const PopupMenuItem(
                              value: 'venue',
                              child: Text('Sort by Venue'),
                            ),
                          ]).then((value) {});
                    },
                  ),
                ],
              ),
            ),
          ),
          SliverPadding(
            padding: const EdgeInsets.all(2),
            sliver: SliverList(
              delegate: SliverChildBuilderDelegate(
                (context, index) {
                  return const EventCard();
                },
                childCount: 10,
              ),
            ),
          ),
        ],
      ),
    );
  }
}
