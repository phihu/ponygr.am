import 'package:flutter/material.dart';

class MessagesPage extends StatelessWidget {
  final ValueChanged<String> navigate;
  const MessagesPage({Key? key, required this.navigate}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: <Widget>[
        Container(
          constraints: const BoxConstraints.expand(),
          child: FittedBox(
            fit: BoxFit.fill,
            child: Container(
              color:Colors.lightGreenAccent,
              child: const Center(
                child: Text('Messages Page'),
              ),
            ),
          ),
        )
      ],
    );
  }
}
