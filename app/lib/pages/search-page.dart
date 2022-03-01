import 'package:flutter/material.dart';
import '../navbar/navbar_widget.dart';

class SearchPage extends StatelessWidget {
  final ValueChanged<String> navigate;
  const SearchPage({Key? key, required this.navigate }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: <Widget>[
          Container(
            constraints: BoxConstraints.expand(),
            child: const Center(
              child: Text('Search Page'),
            ),
          )
        ],
      ),
      bottomNavigationBar:
      BottomNavBar(navigate: this.navigate), // This trailing comma makes auto-formatting nicer for build methods.
    );
  }
}
