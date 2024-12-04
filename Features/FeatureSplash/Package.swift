// swift-tools-version: 6.0

import PackageDescription

let package = Package(
    name: "FeatureSplash",
    platforms: [
        .iOS(.v16),
        .macOS(.v13),
    ],
    products: [
        .library(
            name: "FeatureSplash",
            targets: ["FeatureSplash"]
        ),
    ],
    dependencies: [
        // External
        .package(url: "https://github.com/SimplyDanny/SwiftLintPlugins", exact: "0.57.0"),
        // Core
        .package(path: "../../Core/DesignSystem"),
        .package(path: "../../Core/DesignTheme"),
    ],
    targets: [
        .target(
            name: "FeatureSplash",
            dependencies: [
                .product(name: "DesignSystem", package: "DesignSystem"),
                .product(name: "DesignTheme", package: "DesignTheme"),
            ],
            path: "Sources",
            plugins: [
                .plugin(name: "SwiftLintBuildToolPlugin", package: "SwiftLintPlugins"),
            ]
        ),
    ],
    swiftLanguageModes: [.v6]
)
