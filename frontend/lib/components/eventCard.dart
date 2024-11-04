import 'package:flutter/material.dart';

class EventCard extends StatelessWidget {
  const EventCard({super.key});

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: () {},
      onLongPress: () {
        final RenderBox cardRenderBox = context.findRenderObject() as RenderBox;
        final Offset cardPosition = cardRenderBox.localToGlobal(Offset.zero);
        final RenderBox overlay =
            Overlay.of(context).context.findRenderObject() as RenderBox;

        showMenu(
          context: context,
          color: Colors.white,
          elevation: 8,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(15),
          ),
          position: RelativeRect.fromRect(
            Rect.fromLTWH(
              cardPosition.dx + cardRenderBox.size.width - 16,
              cardPosition.dy + 8,
              0,
              0,
            ),
            Offset.zero & overlay.size,
          ),
          items: [
            const PopupMenuItem(
              value: 'book',
              child: const Row(
                children: const [
                  Icon(Icons.calendar_today, size: 20),
                  SizedBox(width: 12),
                  Text('Book Event', style: TextStyle(fontSize: 16)),
                ],
              ),
            ),
            const PopupMenuItem(
              value: 'share',
              child: const Row(
                children: const [
                  Icon(Icons.share, size: 20),
                  SizedBox(width: 12),
                  Text('Share Event', style: TextStyle(fontSize: 16)),
                ],
              ),
            ),
            const PopupMenuItem(
              value: 'favorite',
              child: const Row(
                children: const [
                  Icon(Icons.favorite_border, size: 20),
                  SizedBox(width: 12),
                  Text('Add to Favorites', style: TextStyle(fontSize: 16)),
                ],
              ),
            ),
          ],
        ).then((value) {
          switch (value) {
            case 'book':
              print('Book event');
              break;
            case 'share':
              print('Share event');
              break;
            case 'favorite':
              print('Add to favorites');
              break;
          }
        });
      },
      child: Container(
        margin: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
        height: 200, // Add a fixed height for the card
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(15),
          boxShadow: [
            BoxShadow(
              color: Colors.grey.withOpacity(0.1),
              spreadRadius: 5,
              blurRadius: 7,
              offset: const Offset(0, 3),
            ),
          ],
        ),
        child: Stack(
          children: [
            // Background image
            ClipRRect(
              borderRadius: BorderRadius.circular(15),
              child: Image.network(
                'https://picsum.photos/500/300',
                width: double.infinity,
                height: double.infinity,
                fit: BoxFit.cover,
              ),
            ),
            // Gradient overlay to make text readable
            Container(
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(15),
                gradient: LinearGradient(
                  begin: Alignment.topCenter,
                  end: Alignment.bottomCenter,
                  colors: [
                    Colors.transparent,
                    Colors.black.withOpacity(0.7),
                  ],
                ),
              ),
            ),
            // Content
            Padding(
              padding: const EdgeInsets.all(16),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                mainAxisAlignment: MainAxisAlignment.end,
                children: [
                  const Text(
                    'Event Name',
                    style: const TextStyle(
                      fontSize: 18,
                      fontWeight: FontWeight.bold,
                      color: Colors.white,
                    ),
                  ),
                  const SizedBox(height: 8),
                  Text(
                    'Event Date',
                    style: TextStyle(
                      fontSize: 16,
                      color: Colors.white.withOpacity(0.9),
                    ),
                  ),
                  const SizedBox(height: 8),
                  Text(
                    'Event Location',
                    style: TextStyle(
                      fontSize: 16,
                      color: Colors.white.withOpacity(0.9),
                    ),
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
