import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class CustomPain extends CustomPainter {
  final Color color;


  CustomPain(this.color);

  @override
  void paint(Canvas canvas, Size size) {
    final curve = Paint()
      ..style = PaintingStyle.stroke
      ..strokeWidth = 3
      ..color = color;

    final path = Path();
    path.moveTo(size.width * .74, 0);
    path.lineTo(size.width * .34, size.height);
    path.lineTo(size.width * .35, size.height);
    path.lineTo(size.width * .75, 0);
    path.close();

    canvas.drawPath(path, curve);
  }

  @override
  bool shouldRepaint(_) {
    return true;
  }
}
