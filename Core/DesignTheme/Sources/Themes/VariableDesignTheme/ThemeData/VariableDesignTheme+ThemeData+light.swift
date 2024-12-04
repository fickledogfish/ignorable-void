extension VariableDesignTheme.ThemeData {
    public static let light = VariableDesignTheme.ThemeData(
        colors: Colors(
            disabled: Colors.Color(
                background: .gray,
                border: .gray,
                placeholderText: .gray,
                text: .gray
            ),
            error: Colors.Color(
                background: .red,
                border: .red,
                placeholderText: .red,
                text: .red
            ),
            normal: Colors.Color(
                background: .white,
                border: .black,
                placeholderText: .gray,
                text: .black
            )
        ),
        spacing: Spacing(
            extraSmall: 3,
            small: 5,
            medium: 10,
            large: 15
        )
    )
}
