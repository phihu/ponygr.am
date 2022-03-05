import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:expandable_text/expandable_text.dart';
import 'package:flutter_glow/flutter_glow.dart';

class PostMessage extends StatelessWidget {
  final String message;

  const PostMessage({Key? key, required this.message}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return  ExpandableText(
      this.message,
      style: TextStyle(
        color: Colors.white,
        fontSize: 14,
        shadows: [
          Shadow(
            color: Colors.black,
            offset: Offset(1.0, 1.0),
            blurRadius: 10.0,
          ),
          // You can add as many Shadow as you want
        ],
      ),
      expandText:  AppLocalizations.of(context)?.postShowMore ?? '...',
      collapseText: AppLocalizations.of(context)?.postShowLess ?? '...',
      maxLines: 5,
      linkColor: Colors.pink,
      animation: true,
      collapseOnTextTap: true,
//      prefixText: username,
//      onPrefixTap: () => showProfile(username),
//      prefixStyle: TextStyle(fontWeight: FontWeight.bold),
//      onHashtagTap: (name) => showHashtag(name),
      hashtagStyle: TextStyle(
        color: Colors.pinkAccent,
        fontWeight: FontWeight.w600,
      ),
//      onMentionTap: (username) => showProfile(username),
      mentionStyle: TextStyle(
        fontWeight: FontWeight.w600,
      ),
//      onUrlTap: (url) => launchUrl(url),
      urlStyle: TextStyle(
        decoration: TextDecoration.underline,
      ),
    );
  }
}
