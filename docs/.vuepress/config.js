module.exports = {
    title: 'VOR Documentation',
    description: 'Welcome to the documentation for VOR',
    base: '/',
    head: [
        ['link', { rel: 'icon', href: '/favicon.ico' }]
    ],
    markdown: {
        // options for markdown-it-toc
        toc: {includeLevel: [2, 3]}
    },

    themeConfig: {
        lastUpdated: 'Last Updated',
        repo: 'unification-com/xfund-vor',
        docsDir: 'docs',
        logo: '/assets/img/unification_logoblack.png',
        sidebar: [
            {
                title: "Introduction",
                path: "/"
            },
            {
                title: "Contract Addresses",
                path: "/contracts"
            },
            {
                title: "VOR Providers",
                path: "/providers"
            },
            {
                title: "Guides",
                path: "/guide",
                children: [
                    "/guide/quickstart",
                    "/guide/implementation",
                    "/guide/interaction",
                    "/guide/advanced",
                    "/guide/oracle",
                ]
            },
            {
                title: "Demos",
                path: "/demos",
                children: [
                    "/demos/nft_demo",
                ]
            },
            {
                title: "Contract Docs",
                children: [
                    "/api/VOR",
                    "/api/VORConsumerBase",
                    "/api/VORCoordinator",
                    "/api/VORRequestIDBase",
                ]
            },
        ],
    }
}
