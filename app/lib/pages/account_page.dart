import 'package:flutter/material.dart';
import '../navbar/navbar_widget.dart';

class AccountPage extends StatelessWidget {
  const AccountPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: <Widget>[
          Container(
            constraints: BoxConstraints.expand(),
            child: Text('ACCOUNT')
          )
        ],
      ),
      bottomNavigationBar:
          const BottomNavBar(), // This trailing comma makes auto-formatting nicer for build methods.
    );
  }
}
