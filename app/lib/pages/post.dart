import 'package:flutter/material.dart';

class PostPage extends StatelessWidget {
  final ValueChanged<String> navigate;
  const PostPage({Key? key, required this.navigate }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Stack(
        children: <Widget>[
          Container(
            constraints: const BoxConstraints.expand(),
            child: const Center(
              child: Text('Post Page'),
            ),
          )
        ],
      );
  }
}
