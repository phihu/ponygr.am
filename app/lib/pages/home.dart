import 'package:flutter/material.dart';

class HomePage extends StatefulWidget {
  final ValueChanged<String> navigate;

  const HomePage({Key? key, required this.navigate }) : super(key: key);

  // This class is the configuration for the state. It holds the values (in this
  // case the title) provided by the parent (in this case the App widget) and
  // used by the build method of the State. Fields in a Widget subclass are
  // always marked "final".

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  @override
  Widget build(BuildContext context) {
    // This method is rerun every time setState is called, for instance as done
    // by the _incrementCounter method above.
    //
    // The Flutter framework has been optimized to make rerunning build methods
    // fast, so that you can just rebuild anything that needs updating rather
    // than having to individually change instances of widgets.
    return Stack(
        children: <Widget>[
          Container(
            constraints: const BoxConstraints.expand(),
            color: Colors.black,
            child: FittedBox(
              fit: BoxFit.fill,
              child: Image.asset('assets/testbg.jpg'),
            ),
          )
        ],
    );
  }
}