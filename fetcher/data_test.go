package fetcher

const (
	GoodBBCResponse = `
		<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet title="XSL_formatting" type="text/xsl" href="/shared/bsp/xsl/rss/nolsol.xsl"?>
<rss xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:atom="http://www.w3.org/2005/Atom" version="2.0" xmlns:media="http://search.yahoo.com/mrss/">
    <channel>
        <title><![CDATA[BBC News - Home]]></title>
        <description><![CDATA[BBC News - Home]]></description>
        <link>https://www.bbc.co.uk/news/</link>
        <image>
            <url>https://news.bbcimg.co.uk/nol/shared/img/bbc_news_120x60.gif</url>
            <title>BBC News - Home</title>
            <link>https://www.bbc.co.uk/news/</link>
        </image>
        <generator>RSS for Node</generator>
        <lastBuildDate>Sat, 30 Jan 2021 14:32:36 GMT</lastBuildDate>
        <copyright><![CDATA[Copyright: (C) British Broadcasting Corporation, see http://news.bbc.co.uk/2/hi/help/rss/4498287.stm for terms and conditions of reuse.]]></copyright>
        <language><![CDATA[en-gb]]></language>
        <ttl>15</ttl>
        <item>
            <title><![CDATA[EU vaccine export row: Bloc backtracks on controls for NI]]></title>
            <description><![CDATA[It follows a decision to invoke an emergency provision in the Brexit deal in order to control vaccine exports.]]></description>
            <link>https://www.bbc.co.uk/news/uk-55865539</link>
            <guid isPermaLink="true">https://www.bbc.co.uk/news/uk-55865539</guid>
            <pubDate>Sat, 30 Jan 2021 14:15:18 GMT</pubDate>
        </item>
        <item>
            <title><![CDATA[Arlene Foster urges PM to replace 'unworkable' NI Brexit deal]]></title>
            <description><![CDATA[Arlene Foster wants GB-NI trade flow problems addressed, in the wake of the EU vaccine export row.]]></description>
            <link>https://www.bbc.co.uk/news/uk-northern-ireland-55866285</link>
            <guid isPermaLink="true">https://www.bbc.co.uk/news/uk-northern-ireland-55866285</guid>
            <pubDate>Sat, 30 Jan 2021 14:21:14 GMT</pubDate>
        </item>
    </channel>
</rss>
`

	OldVersionBBCResponse = `
		<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet title="XSL_formatting" type="text/xsl" href="/shared/bsp/xsl/rss/nolsol.xsl"?>
<rss xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:atom="http://www.w3.org/2005/Atom" version="1.0" xmlns:media="http://search.yahoo.com/mrss/">
    <channel>
        <title><![CDATA[BBC News - Home]]></title>
        <description><![CDATA[BBC News - Home]]></description>
        <link>https://www.bbc.co.uk/news/</link>
        <image>
            <url>https://news.bbcimg.co.uk/nol/shared/img/bbc_news_120x60.gif</url>
            <title>BBC News - Home</title>
            <link>https://www.bbc.co.uk/news/</link>
        </image>
        <generator>RSS for Node</generator>
        <lastBuildDate>Sat, 30 Jan 2021 14:32:36 GMT</lastBuildDate>
        <copyright><![CDATA[Copyright: (C) British Broadcasting Corporation, see http://news.bbc.co.uk/2/hi/help/rss/4498287.stm for terms and conditions of reuse.]]></copyright>
        <language><![CDATA[en-gb]]></language>
        <ttl>15</ttl>
        <item>
            <title><![CDATA[EU vaccine export row: Bloc backtracks on controls for NI]]></title>
            <description><![CDATA[It follows a decision to invoke an emergency provision in the Brexit deal in order to control vaccine exports.]]></description>
            <link>https://www.bbc.co.uk/news/uk-55865539</link>
            <guid isPermaLink="true">https://www.bbc.co.uk/news/uk-55865539</guid>
            <pubDate>Sat, 30 Jan 2021 14:15:18 GMT</pubDate>
        </item>
        <item>
            <title><![CDATA[Arlene Foster urges PM to replace 'unworkable' NI Brexit deal]]></title>
            <description><![CDATA[Arlene Foster wants GB-NI trade flow problems addressed, in the wake of the EU vaccine export row.]]></description>
            <link>https://www.bbc.co.uk/news/uk-northern-ireland-55866285</link>
            <guid isPermaLink="true">https://www.bbc.co.uk/news/uk-northern-ireland-55866285</guid>
            <pubDate>Sat, 30 Jan 2021 14:21:14 GMT</pubDate>
        </item>
    </channel>
</rss>
`

	MissingFieldsBBCResponse = `
		<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet title="XSL_formatting" type="text/xsl" href="/shared/bsp/xsl/rss/nolsol.xsl"?>
<rss xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:atom="http://www.w3.org/2005/Atom" version="2.0" xmlns:media="http://search.yahoo.com/mrss/">
    <channel>
        <title><![CDATA[BBC News - Home]]></title>
        <description><![CDATA[BBC News - Home]]></description>
        <link>https://www.bbc.co.uk/news/</link>
        <image>
            <url>https://news.bbcimg.co.uk/nol/shared/img/bbc_news_120x60.gif</url>
            <title>BBC News - Home</title>
            <link>https://www.bbc.co.uk/news/</link>
        </image>
        <generator>RSS for Node</generator>
        <lastBuildDate>Sat, 30 Jan 2021 14:32:36 GMT</lastBuildDate>
        <copyright><![CDATA[Copyright: (C) British Broadcasting Corporation, see http://news.bbc.co.uk/2/hi/help/rss/4498287.stm for terms and conditions of reuse.]]></copyright>
        <language><![CDATA[en-gb]]></language>
        <ttl>15</ttl>
        <item>
            <title><![CDATA[EU vaccine export row: Bloc backtracks on controls for NI]]></title>
            <description><![CDATA[It follows a decision to invoke an emergency provision in the Brexit deal in order to control vaccine exports.]]></description>
            <link>https://www.bbc.co.uk/news/uk-55865539</link>
        </item>
        <item>
        </item>
        <item>
            <description><![CDATA[Arlene Foster wants GB-NI trade flow problems addressed, in the wake of the EU vaccine export row.]]></description>
            <link>https://www.bbc.co.uk/news/uk-northern-ireland-55866285</link>
        </item>
        <item>
            <title><![CDATA[EU vaccine export row: Bloc backtracks on controls for NI]]></title>
            <link>https://www.bbc.co.uk/news/uk-northern-ireland-55866285</link>
        </item>
        <item>
            <title><![CDATA[EU vaccine export row: Bloc backtracks on controls for NI]]></title>
            <description><![CDATA[Arlene Foster wants GB-NI trade flow problems addressed, in the wake of the EU vaccine export row.]]></description>
        </item>
    </channel>
</rss>
`

	InvalidChannelBBCResponse = `
		<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet title="XSL_formatting" type="text/xsl" href="/shared/bsp/xsl/rss/nolsol.xsl"?>
<rss xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:atom="http://www.w3.org/2005/Atom" version="2.0" xmlns:media="http://search.yahoo.com/mrss/">
    <channel>
        <title><![CDATA[BBC News - Home]]></title>
        <description><![CDATA[BBC News - Home]]></description>
        <image>
            <url>https://news.bbcimg.co.uk/nol/shared/img/bbc_news_120x60.gif</url>
            <title>BBC News - Home</title>
            <link>https://www.bbc.co.uk/news/</link>
        </image>
        <generator>RSS for Node</generator>
        <lastBuildDate>Sat, 30 Jan 2021 14:32:36 GMT</lastBuildDate>
        <copyright><![CDATA[Copyright: (C) British Broadcasting Corporation, see http://news.bbc.co.uk/2/hi/help/rss/4498287.stm for terms and conditions of reuse.]]></copyright>
        <language><![CDATA[en-gb]]></language>
        <ttl>15</ttl>
        <item>
            <title><![CDATA[EU vaccine export row: Bloc backtracks on controls for NI]]></title>
            <description><![CDATA[It follows a decision to invoke an emergency provision in the Brexit deal in order to control vaccine exports.]]></description>
            <link>https://www.bbc.co.uk/news/uk-55865539</link>
        </item>
    </channel>
</rss>
`

	InvalidXMLBBCResponse = `
		<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet title="XSL_formatting" type="text/xsl" href="/shared/bsp/xsl/rss/nolsol.xsl"?>
<rss xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:atom="http://www.w3.org/2005/Atom" version="2.0" xmlns:media="http://search.yahoo.com/mrss/">
        <title><![CDATA[BBC News - Home]]></title>
        <description><![CDATA[BBC News - Home]]></description>
        <image>
            <url>https://news.bbcimg.co.uk/nol/shared/img/bbc_news_120x60.gif</url>
            <title>BBC News - Home</title>
            <link>https://www.bbc.co.uk/news/</link>
        </image>
        <generator>RSS for Node</generator>
        <lastBuildDate>Sat, 30 Jan 2021 14:32:36 GMT</lastBuildDate>
        <copyright><![CDATA[Copyright: (C) British Broadcasting Corporation, see http://news.bbc.co.uk/2/hi/help/rss/4498287.stm for terms and conditions of reuse.]]></copyright>
        <language><![CDATA[en-gb]]></language>
        <ttl>15</ttl>
        <item>
            <title><![CDATA[EU vaccine export row: Bloc backtracks on controls for NI]]></title>
            <description><![CDATA[It follows a decision to invoke an emergency provision in the Brexit deal in order to control vaccine exports.]]></description>
            <link>https://www.bbc.co.uk/news/uk-55865539</link>
        </item>
    </channel>
</rss>
`
	EmptyXMLBBCResponse = `
		<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet title="XSL_formatting" type="text/xsl" href="/shared/bsp/xsl/rss/nolsol.xsl"?>
<rss xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:atom="http://www.w3.org/2005/Atom" version="2.0" xmlns:media="http://search.yahoo.com/mrss/">
</rss>
`
)
