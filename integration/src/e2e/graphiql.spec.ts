import {expect, test} from '@playwright/test';

const uri = process.env.VITE_SERVER_URL || 'http://localhost:8080';

test('has title', async ({page}) => {
    await page.goto(uri);

    await expect(page).toHaveTitle("GraphQL playground");
});

test('has execute button', async ({page}) => {
    await page.goto(uri);

    const executeButton = await page.locator('button[class=graphiql-execute-button]');
    await expect(executeButton).toBeDefined();
    await expect(await executeButton.count()).toBe(1);
});

test('has graphiql editor', async ({page}) => {
    await page.goto(uri);

    const selector = 'div[class=graphiql-editor]';
    await page.waitForSelector(selector);
    const editorDiv = await page.locator(selector);
    await expect(editorDiv).toBeDefined();
    await expect(await editorDiv.count()).toBe(2);
});
