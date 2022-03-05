import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';

class PostWidget extends StatefulWidget {
  const PostWidget({Key? key}) : super(key: key);

  @override
  _PostWidgetState createState() => _PostWidgetState();
}

class _PostWidgetState extends State<PostWidget> {
  static const iconBoxSize = 80.0;
  static const iconSize = 55.0;

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        Container(
          constraints: const BoxConstraints.expand(),
          color: Colors.black,
          child: FittedBox(
            fit: BoxFit.fill,
            child: Image.asset('assets/testbg.jpg'),
          ),
        ),
        SafeArea(
          bottom:false,
          child: Row(
            mainAxisSize: MainAxisSize.max,
            mainAxisAlignment: MainAxisAlignment.end,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Column(
                mainAxisAlignment: MainAxisAlignment.end,
                children: const [
                  SizedBox(
                    height: iconBoxSize,
                    child: FaIcon(FontAwesomeIcons.solidHeart, size: iconSize, color: Colors.white),
                  ),
                  SizedBox(
                      height: iconBoxSize,
                      child: FaIcon(FontAwesomeIcons.icons, size: iconSize, color: Colors.white),
                  ),
                  SizedBox(
                    height: iconBoxSize,
                    child:
                      FaIcon(FontAwesomeIcons.reply, size: iconSize, color: Colors.white),
                  ),
                ]
              ),
              Expanded(
                child: Container(
                  decoration: BoxDecoration(
                      border: Border.all(color: Colors.redAccent)
                  ),
                )
              )
            ]
          ),
        )
      ],
    );
  }
}

