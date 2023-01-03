import SunEditor from 'suneditor-react';
import 'suneditor/dist/css/suneditor.min.css';
import katex from 'katex';
import 'katex/dist/katex.min.css';
import { useState } from 'react';

function FormSoalComponent(props) {
    const btnDefault = [
        // Default
        ['undo', 'redo'],
        ['font', 'fontSize', 'formatBlock'],
        ['paragraphStyle', 'blockquote'],
        ['bold', 'underline', 'italic', 'strike', 'subscript', 'superscript'],
        ['fontColor', 'hiliteColor', 'textStyle'],
        ['removeFormat'],
        ['outdent', 'indent'],
        ['align', 'horizontalRule', 'list', 'lineHeight'],
        ['table', 'link', 'image', 'video', 'audio', 'math'],
        ['imageGallery'],
        ['fullScreen', 'showBlocks', 'codeView'],
        ['preview', 'print'],
        ['save', 'template'],
        ['-left', '#fix', 'dir_ltr', 'dir_rtl'],
        // (min-width:992px)
        ['%992', [
            ['undo', 'redo'],
            [':p-More Paragraph-default.more_paragraph', 'font', 'fontSize', 'formatBlock', 'paragraphStyle', 'blockquote'],
            ['bold', 'underline', 'italic', 'strike'],
            [':t-More Text-default.more_text', 'subscript', 'superscript', 'fontColor', 'hiliteColor', 'textStyle'],
            ['removeFormat'],
            ['outdent', 'indent'],
            ['align', 'horizontalRule', 'list', 'lineHeight'],
            ['-right', 'dir'],
            ['-right', ':i-More Misc-default.more_vertical', 'fullScreen', 'showBlocks', 'codeView', 'preview', 'print', 'save', 'template'],
            ['-right', ':r-More Rich-default.more_plus', 'table', 'link', 'image', 'video', 'audio', 'math', 'imageGallery']
        ]],
        // (min-width:768px)
        ['%768', [
            ['undo', 'redo'],
            [':p-More Paragraph-default.more_paragraph', 'font', 'fontSize', 'formatBlock', 'paragraphStyle', 'blockquote'],
            [':t-More Text-default.more_text', 'bold', 'underline', 'italic', 'strike', 'subscript', 'superscript', 'fontColor', 'hiliteColor', 'textStyle', 'removeFormat'],
            [':e-More Line-default.more_horizontal', 'outdent', 'indent', 'align', 'horizontalRule', 'list', 'lineHeight'],
            [':r-More Rich-default.more_plus', 'table', 'link', 'image', 'video', 'audio', 'math', 'imageGallery'],
            ['-right', 'dir'],
            ['-right', ':i-More Misc-default.more_vertical', 'fullScreen', 'showBlocks', 'codeView', 'preview', 'print', 'save', 'template']
        ]]
    ];

    let [content, setContent] = useState('');
    let [hide, setHide] = useState(false);

    const handleChange = (content) => {
        setContent(content);
    }

    const handleClick = () => {
        hide ? setHide(false) : setHide(true);
    }

    // useEffect(() => {
    //     fetch('http://localhost:3000/auth/lists')
    //         .then(response => response.json())
    //         .then(obj => console.log(obj))
    //         .catch((err) => console.log(err))
    // })

    return (
        <>
            <SunEditor
                setOptions={{
                    height: 200,
                    buttonList: btnDefault,
                    katex: katex,
                }}
                onChange={handleChange}
                hide={hide}
            />
            {
                hide ?
                <>
                    <button onClick={handleClick}>Edit Question</button>
                    <div dangerouslySetInnerHTML={{ __html: content }}>
                    </div>
                </>
                :
                <>
                    <button onClick={handleClick}>Submit</button>
                    <div></div>
                </>
            }
        </>
    );
}

export default FormSoalComponent;