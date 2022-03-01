import 'package:flutter/material.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'navbar/navbar-widget.dart';

void main() {
  runApp(const PonygramApp());
}

class PonygramApp extends StatelessWidget {
  const PonygramApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Ponygr.am',
      localizationsDelegates: [
        AppLocalizations.delegate, // Add this line
        GlobalMaterialLocalizations.delegate,
        GlobalWidgetsLocalizations.delegate,
        GlobalCupertinoLocalizations.delegate,
      ],
      supportedLocales: [
        Locale('en', ''), // English, no country code
        Locale('es', ''), // Spanish, no country code
      ],
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: const PonygramHomePage(title: 'title'),
    );
  }
}

class PonygramHomePage extends StatefulWidget {
  const PonygramHomePage({Key? key, required this.title}) : super(key: key);

  // This class is the configuration for the state. It holds the values (in this
  // case the title) provided by the parent (in this case the App widget) and
  // used by the build method of the State. Fields in a Widget subclass are
  // always marked "final".

  final String title;

  @override
  State<PonygramHomePage> createState() => _PonygramPageState();
}

class _PonygramPageState extends State<PonygramHomePage> {
  int _counter = 0;

  void _incrementCounter() {
    setState(() {
      // This call to setState tells the Flutter framework that something has
      // changed in this State, which causes it to rerun the build method below
      // so that the display can reflect the updated values. If we changed
      // _counter without calling setState(), then the build method would not be
      // called again, and so nothing would appear to happen.
      _counter++;
    });
  }


  @override
  Widget build(BuildContext context) {
    // This method is rerun every time setState is called, for instance as done
    // by the _incrementCounter method above.
    //
    // The Flutter framework has been optimized to make rerunning build methods
    // fast, so that you can just rebuild anything that needs updating rather
    // than having to individually change instances of widgets.
    return Scaffold(
      body: Stack(
        children:<Widget>[
          Container(
            constraints: BoxConstraints.expand(),
            color:Colors.black,
            child:
              FittedBox(
                  fit: BoxFit.fill,
                  child: Image.asset('assets/testbg.jpg'),
              ),
          )
    ],
      ),
      bottomNavigationBar: const BottomNavBar(), // This trailing comma makes auto-formatting nicer for build methods.
    );
  }
}
