// swift-tools-version: 6.0

import PackageDescription

let package = Package(
    name: "Network",
    platforms: [
        .iOS(.v16),
        .macOS(.v13),
    ],
    products: [
        .library(
            name: "Network",
            targets: ["Http"]
        ),
    ],
    dependencies: [
        // External
        .package(
            url: "https://github.com/SimplyDanny/SwiftLintPlugins",
            exact: "0.57.0"
        ),
        .package(
            url: "https://github.com/Alamofire/Alamofire",
            exact: "5.10.0"
        ),
    ],
    targets: [
        .target(
            name: "Http",
            dependencies: [
                .product(name: "Alamofire", package: "Alamofire"),
            ],
            path: "Sources/Http",
            plugins: [
                .plugin(name: "SwiftLintBuildToolPlugin", package: "SwiftLintPlugins"),
            ]
        ),
        .testTarget(
            name: "HttpTests",
            dependencies: [
                .target(name: "Http"),
            ],
            path: "Tests/Http"
        ),
    ]
)
