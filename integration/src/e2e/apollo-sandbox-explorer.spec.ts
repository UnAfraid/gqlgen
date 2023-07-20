import {expect, test} from '@playwright/test';

const uri = `${process.env.VITE_SERVER_URL || 'http://localhost:8080'}/apollo`;

test('has title', async ({page}) => {
    await page.goto(uri);

    await expect(page).toHaveTitle("GraphQL playground");
});

test('has frame', async ({page}) => {
    await page.goto(uri);

    const selector = "#apollo-embed-0"

    await page.waitForSelector(selector);
    const frame = await page.frameLocator(selector);
    await expect(frame).toBeDefined()
});
