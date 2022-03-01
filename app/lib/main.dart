import 'package:flutter/material.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

import 'routing/ponygram_route_information_parser.dart';
import 'routing/ponygram_router_delegate.dart';

import 'pages/home_page.dart';
import 'pages/account_page.dart';

void main() {
  runApp(const PonygramApp());
}

class PonygramApp extends StatefulWidget {
  const PonygramApp({Key? key}) : super(key: key);

  @override
  _PonygramAppState createState() => _PonygramAppState();
}

class _PonygramAppState extends State<PonygramApp> {
  PonygramRouterDelegate _routerDelegate = PonygramRouterDelegate();
  PonygramRouteInformationParser _routeInformationParser =
  PonygramRouteInformationParser();

  @override

  Widget build(BuildContext context) {
    return MaterialApp.router(
      title: 'Ponygr.am',
      routerDelegate: _routerDelegate,
      routeInformationParser: _routeInformationParser,
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
      )
    );
  }
}


