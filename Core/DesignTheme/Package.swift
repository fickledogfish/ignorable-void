// swift-tools-version: 6.0

import PackageDescription

let package = Package(
    name: "Theme",
    platforms: [
        .iOS(.v16),
        .macOS(.v13)
    ],
    products: [
        .library(
            name: "DesignTheme",
            targets: ["DesignTheme"]
        ),
    ],
    dependencies: [
        // External
        .package(url: "https://github.com/SimplyDanny/SwiftLintPlugins", exact: "0.57.0"),
    ],
    targets: [
        .target(
            name: "DesignTheme",
            path: "Sources",
            plugins: [
                .plugin(name: "SwiftLintBuildToolPlugin", package: "SwiftLintPlugins"),
            ]
        ),
        .testTarget(
            name: "DesignThemeTests",
            dependencies: [
                .target(name: "DesignTheme"),
            ],
            path: "Tests"
        ),
    ],
    swiftLanguageModes: [.v6]
)
