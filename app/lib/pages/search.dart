import 'package:flutter/material.dart';

class SearchPage extends StatelessWidget {
  final ValueChanged<String> navigate;
  const SearchPage({Key? key, required this.navigate }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Stack(
        children: <Widget>[
          Container(
            constraints: const BoxConstraints.expand(),
            child: const Center(
              child: Text('Search Page'),
            ),
          )
        ],
      );
  }
}
