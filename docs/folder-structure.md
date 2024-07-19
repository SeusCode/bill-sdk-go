# Bill - Folder structure

```
bill-sdk-go/
├── api/
│   ├── afip/
│   ├── ...
├── config/
├── docs/
├── examples/
├── pkg/
│   ├── http/
│   ├── endpoints/
├── models/
│   ├── afip/
│   │   ├── ...
│   ├── api/
├── go.mod
├── go.sum
└── README.md
```

## Folder Explanation

-   api/: This directory contains all the logic for communicating with different parts of your API. It is organized into subdirectories for each module of the API.

-   config/: Contains configuration management code.

-   docs/: Contains documentation files.

-   examples/: Contains example code to demonstrate how to use the SDK.

-   internal/: Contains internal packages that are not meant to be used by external code.

-   pkg/: Contains packages that provide utility functions and shared components that can be used across the SDK.

This structure ensures that the code is modular, organized, and maintainable, making it easy to add new features, fix bugs, and scale the project over time.
