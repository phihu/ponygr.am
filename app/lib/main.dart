import 'package:flutter/material.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:ponygram/widgets/navbar_widget.dart';

import 'package:ponygram/l10n/localization.dart';
import 'package:ponygram/routing/pny_route_information_parser.dart';
import 'package:ponygram/routing/ponygram_router_delegate.dart';

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
      localeResolutionCallback: (deviceLocale, supportedLocales) {
        if (supportedLocales
            .map((e) => e.languageCode)
            .contains(deviceLocale?.languageCode)) {
          return deviceLocale;
        } else {
          return const Locale('en', '');
        }
      },
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      builder: (context, child){
        var t = AppLocalizations.of(context);
        // Initialize localization with main context
        Localization.init(context);
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


