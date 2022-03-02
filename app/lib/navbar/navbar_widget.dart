import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class BottomNavBar extends StatefulWidget {
  final ValueChanged<String> navigate;
  const BottomNavBar({Key? key, required this.navigate}) : super(key: key);

  @override
  _BottomNavBarState createState() => _BottomNavBarState();
}

class _BottomNavBarState extends State<BottomNavBar> {
  int _selectedIndex = 0;

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
      switch(index){
        case 1:
          widget.navigate('search');
          break;
        case 2:
          widget.navigate('post');
          break;
        case 3:
          widget.navigate('messages');
          break;
        case 4:
          widget.navigate('account');
          break;
        default:
          widget.navigate('home');
          break;
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return BottomNavigationBar(
      showSelectedLabels: false,
      showUnselectedLabels: false,
      selectedFontSize: 0,
      backgroundColor: Colors.black,
      type: BottomNavigationBarType.fixed,
      items: <BottomNavigationBarItem>[
        BottomNavigationBarItem(
          icon: Icon(Icons.home_rounded, size: 30.0, color: Colors.purple),
          label: AppLocalizations.of(context)?.navLabelHome ?? '',
          backgroundColor: Colors.red,
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.search_rounded, size: 30.0, color: Colors.yellowAccent),
          label: AppLocalizations.of(context)?.navLabelSearch ?? '',
          backgroundColor: Colors.green,
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.add_circle_rounded, size: 70.0, color: Colors.cyanAccent),
          label: AppLocalizations.of(context)?.navLabelNewPost ?? '',
          backgroundColor: Colors.purple,
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.mail_rounded, size: 30.0, color: Colors.pink),
          label: AppLocalizations.of(context)?.navLabelMessages ?? '',
          backgroundColor: Colors.purple,
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.person_rounded, size: 30.0, color: Colors.greenAccent),
          label: AppLocalizations.of(context)?.navLabelAccount ?? '',
          backgroundColor: Colors.pink,
        ),
      ],
      currentIndex: _selectedIndex,
      selectedItemColor: Colors.amber[800],
      onTap: _onItemTapped,
    );
  }
}
