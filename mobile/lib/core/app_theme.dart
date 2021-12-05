import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

final ThemeData appThemeData = ThemeData(
  primaryColor: Colors.white,
  // accentColor: AppColors.assets,
  appBarTheme: const AppBarTheme(
    elevation: 0,
    color: Colors.white,
    // actionsIconTheme: IconThemeData(color: Colors.black),
    centerTitle: true,
    iconTheme: IconThemeData(color: first),
    systemOverlayStyle: SystemUiOverlayStyle(
      statusBarColor: Color(0xFFff8729),
      statusBarIconBrightness: Brightness.dark,
    ),
  ),
  bottomSheetTheme: const BottomSheetThemeData(backgroundColor: Colors.transparent),
  scaffoldBackgroundColor: Colors.white,
  // fontFamily: 'Georgia',
  textTheme: const TextTheme(
    headline1: TextStyle(fontSize: 72.0, fontWeight: FontWeight.bold),
  ),
);

const Color first = Color(0xFFea3435);
const  Color second = Color(0xFFff8729);