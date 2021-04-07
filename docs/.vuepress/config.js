module.exports = {
    title: 'VOR Documentation',
    description: 'Welcome to the documentation for VOR',
    base: '/',
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
            },{
                title: "Contract Addresses",
                path: "/contracts"
            },
            {
                title: "Guides",
                path: "/guide",
                children: [
                    "/guide/quickstart",
                ]
            },
            {
                title: "Contract Docs",
                children: [
                    "/api/Router",
                    "/api/lib/ConsumerLib",
                    "/api/lib/ConsumerBase",
                ]
            },
        ],
    }
}
