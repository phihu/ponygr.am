import 'package:flutter/material.dart';
import '../navbar/navbar_widget.dart';

class MessagesPage extends StatelessWidget {
  final ValueChanged<String> navigate;
  const MessagesPage({Key? key, required this.navigate }) : super(key: key);

  @override
  Widget build(BuildContext context) {
      return Stack(
        children: <Widget>[
          Container(
            constraints: BoxConstraints.expand(),
            child: const Center(
              child: Text('Messages Page'),
            ),
          )
        ],
      );
  }
}
