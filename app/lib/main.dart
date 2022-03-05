import 'package:flutter/material.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import '../navbar/navbar_widget.dart';

import 'routing/pny_route_information_parser.dart';
import 'routing/ponygram_router_delegate.dart';

void main() {
  runApp(const PonygramApp());
}

class PonygramApp extends StatefulWidget {
  const PonygramApp({Key? key}) : super(key: key);

  @override
  _PonygramAppState createState() => _PonygramAppState();
}

class _PonygramAppState extends State<PonygramApp> {
  final PonygramRouterDelegate _routerDelegate = PonygramRouterDelegate();
  final  PonygramRouteInformationParser _routeInformationParser =
  PonygramRouteInformationParser();

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      title: 'Ponygr.am',
      routerDelegate: _routerDelegate,
      routeInformationParser: _routeInformationParser,
      localizationsDelegates: const [
        AppLocalizations.delegate,
        GlobalMaterialLocalizations.delegate,
        GlobalWidgetsLocalizations.delegate,
        GlobalCupertinoLocalizations.delegate,
      ],
      supportedLocales: const [
        Locale('en', ''), // English, no country code
        Locale('es', ''), // Spanish, no country code
      ],
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      builder: (context, child){
        return MaterialApp(
            home: Scaffold(
              body:child,
              bottomNavigationBar:
                BottomNavBar(navigate: _routerDelegate.navigate), // This trailing comma makes auto-formatting nicer for build methods.
            ),
        );
      }
    );
  }
}


