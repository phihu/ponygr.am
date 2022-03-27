import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:ponygram/l10n/localization.dart';

class BottomNavBar extends StatefulWidget{
  final ValueChanged<String> navigate;
  const BottomNavBar({Key? key, required this.navigate}) : super(key: key);

  @override
  _BottomNavBarState createState() => _BottomNavBarState();
}

class _BottomNavBarState extends State<BottomNavBar> with Localization {
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
          icon: const FaIcon(FontAwesomeIcons.home, size: 30.0, color: Colors.purple),
          label: localize?.navLabelHome ?? '',
          backgroundColor: Colors.red,
        ),
        BottomNavigationBarItem(
          icon: const FaIcon(FontAwesomeIcons.search, size: 30.0, color: Colors.yellowAccent),
          label: localize?.navLabelSearch ?? '',
          backgroundColor: Colors.green,
        ),
        BottomNavigationBarItem(
          icon: const FaIcon(FontAwesomeIcons.plusCircle, size: 70.0, color: Colors.cyanAccent),
          label: localize?.navLabelNewPost ?? '',
          backgroundColor: Colors.purple,
        ),
        BottomNavigationBarItem(
          icon: const FaIcon(FontAwesomeIcons.envelope, size: 30.0, color: Colors.pink),
          label: localize?.navLabelMessages ?? '',
          backgroundColor: Colors.purple,
        ),
        BottomNavigationBarItem(
          icon: const FaIcon(FontAwesomeIcons.userCog, size: 30.0, color: Colors.greenAccent),
          label: localize?.navLabelAccount ?? '',
          backgroundColor: Colors.pink,
        ),
      ],
      currentIndex: _selectedIndex,
      selectedItemColor: Colors.amber[800],
      onTap: _onItemTapped,
    );
  }
}
