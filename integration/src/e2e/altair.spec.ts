import {expect, test} from '@playwright/test';

const uri = `${process.env.VITE_SERVER_URL || 'http://localhost:8080'}/altair`;

test('has title', async ({page}) => {
    await page.goto(uri);

    await expect(page).toHaveTitle("GraphQL playground");
});

test('has send request button', async ({page}) => {
    await page.goto(uri);

    const selector = 'button[track-id="set_request"]';
    await page.waitForSelector(selector);
    const executeButton = await page.locator(selector);
    await expect(executeButton).toBeDefined();
    await expect(await executeButton.count()).toBe(1);
    await expect((await executeButton.nth(0).textContent())?.trim()).toBe('Send Request');
});
